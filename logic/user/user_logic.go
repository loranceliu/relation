package user

import (
	"fmt"
	"gin/entity/dto"
	"gin/entity/vo"
	"gin/errors"
	"gin/model"
	"gin/svc"
	"gin/types/relation"
	"gin/types/user"
	"gin/utils"
	"github.com/jinzhu/copier"
	"strconv"
)

func Login(ctx *svc.ServiceContext, req *user.UserRequest) (usVo *vo.UserInfoVO, err error) {
	var userInfo vo.UserInfoVO
	var sysUser model.SysUser
	er := &errors.ForbiddenError{Message: "用户名或密码错误"}

	if err := model.DB.Raw("select user_id,is_admin,username,password,salt,name,email from tb_sys_user where status = 1 and (username = ? or email = ?)", req.Username, req.Username).Scan(&sysUser).Error; err != nil {
		return nil, er
	}

	if sysUser.UserID == 0 {
		return nil, er
	}

	pass := utils.HashMD5(req.Password + sysUser.Salt)

	if sysUser.Password != pass {
		return nil, er
	}

	err = copier.Copy(&userInfo, &sysUser)
	if err != nil {
		return nil, er
	}

	token, exp, err := utils.GenerateToken(&userInfo)
	if err != nil {
		return nil, er
	}

	userInfo.Token = token
	userInfo.Exp = exp
	userInfo.IsAdmin = sysUser.IsAdmin == 2

	return &userInfo, nil
}

func GetUserInfo(ctx *svc.ServiceContext) (usVo *vo.UserInfoVO, err error) {
	var userInfo vo.UserInfoVO
	var sysUser model.SysUser
	er := &errors.ForbiddenError{Message: "无此用户信息"}

	userId := ctx.Value("user_id")

	if err := model.DB.Raw("select user_id,is_admin,username,password,salt,name,email from tb_sys_user where status = 1 and user_id = ?", userId).Scan(&sysUser).Error; err != nil {
		return nil, er
	}

	if sysUser.UserID == 0 {
		return nil, er
	}

	err = copier.Copy(&userInfo, &sysUser)
	if err != nil {
		return nil, er
	}

	userInfo.IsAdmin = sysUser.IsAdmin == 2

	return &userInfo, nil
}

func GetUserMineInfo(ctx *svc.ServiceContext) (usVo *vo.UserMineInfoVO, err error) {
	var userMineInfo vo.UserMineInfoVO
	var sysUser model.SysUser
	var userStatistics []dto.UserStatisticsDto
	er := &errors.ForbiddenError{Message: "无此用户信息"}

	userId := ctx.Value("user_id")

	if err := model.DB.Raw("select user_id,is_admin,username,password,salt,name,email from tb_sys_user where status = 1 and user_id = ?", userId).Scan(&sysUser).Error; err != nil {
		return nil, er
	}

	if err := model.DB.Raw("select transaction_type type ,sum(money) value,count(transaction_type) num  from tb_relation where owner_id = ? GROUP BY transaction_type", userId).Scan(&userStatistics).Error; err != nil {
		return nil, er
	}

	if err := model.DB.Raw("SELECT COUNT(*) AS days FROM ( SELECT DISTINCT DATE(create_time) AS date FROM tb_relation WHERE owner_id = ?) AS days_with_data", userId).Scan(&userMineInfo.Time).Error; err != nil {
		return nil, er
	}

	if sysUser.UserID == 0 {
		return nil, er
	}

	var v1 float32
	var v2 float32

	if len(userStatistics) > 0 {
		for _, s := range userStatistics {
			if s.Type == 1 {
				userMineInfo.IncomeNum = s.Num
				v1 = s.Value
			}

			if s.Type == 2 {
				userMineInfo.ExpendNum = s.Num
				v2 = s.Value
			}

		}
	}

	userMineInfo.UserId = sysUser.UserID
	userMineInfo.Name = sysUser.Name
	userMineInfo.Email = sysUser.Email
	userMineInfo.Revenue = v1 - v2
	return &userMineInfo, nil
}

func AddUser(ctx *svc.ServiceContext, req *user.UserRequest) (inter interface{}, err error) {
	rl := model.SysUser{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	salt := utils.GenerateSalt()

	// 将密码和盐值组合
	passwordWithSalt := req.Password + salt

	// 计算 MD5 哈希值
	hashedPassword := utils.HashMD5(passwordWithSalt)

	rl.Salt = salt
	rl.Password = hashedPassword

	model.DB.Create(&rl)
	return
}

func UpdateUser(ctx *svc.ServiceContext, req *user.UserRequest) (inter interface{}, err error) {
	rl := model.SysUser{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	salt := utils.GenerateSalt()

	if req.Password != "" {
		// 将密码和盐值组合
		passwordWithSalt := req.Password + salt

		// 计算 MD5 哈希值
		hashedPassword := utils.HashMD5(passwordWithSalt)

		rl.Salt = salt
		rl.Password = hashedPassword
	}

	model.DB.Updates(&rl)
	return
}

func UpdatePasswordUser(ctx *svc.ServiceContext, req *user.UserRequest) (inter interface{}, err error) {
	rl := model.SysUser{}

	if err != nil {
		return nil, err
	}
	salt := utils.GenerateSalt()

	if req.Password != "" {
		// 将密码和盐值组合
		passwordWithSalt := req.Password + salt

		// 计算 MD5 哈希值
		hashedPassword := utils.HashMD5(passwordWithSalt)

		rl.Salt = salt
		rl.Password = hashedPassword
	}

	userId := ctx.Value("user_id")

	userId64, err := strconv.ParseUint(userId, 10, 64)

	rl.UserID = uint(userId64)

	updateFields := map[string]interface{}{
		"password": rl.Password,
		"salt":     rl.Salt,
	}

	model.DB.Model(&rl).Updates(updateFields)
	return
}

func UpdatePersonalUser(ctx *svc.ServiceContext, req *user.UserRequest) (inter interface{}, err error) {
	rl := model.SysUser{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}

	userId := ctx.Value("user_id")

	userId64, err := strconv.ParseUint(userId, 10, 64)

	rl.UserID = uint(userId64)

	updateFields := map[string]interface{}{
		"name":  rl.Name,
		"email": rl.Email,
	}

	model.DB.Model(&rl).Updates(updateFields)
	return
}

func GetSystemUserList(ctx *svc.ServiceContext, req *relation.SystemUserPageRequest) (*vo.PageVO, error) {
	// 构建原生 SQL 查询语句
	sqlQuery := "select user_id, username, email, name, status from tb_sys_user where 1 = 1"
	var sqlArgs []interface{}

	// 添加过滤条件
	if req.Name != "" {
		sqlQuery += " AND name like ?"
		sqlArgs = append(sqlArgs, fmt.Sprintf("%%%s%%", req.Name))
	}

	if req.Status != 0 {
		sqlQuery += " AND status = ?"
		sqlArgs = append(sqlArgs, req.Status)
	}

	sqlQuery += " Order By user_id desc"

	page := utils.New(sqlQuery, sqlArgs, &req.PageRequest)
	sql := page.StartPage()

	var systemUsers []vo.SystemUserVO
	if err := model.DB.Raw(sql, sqlArgs...).Scan(&systemUsers).Error; err != nil {
		return nil, err
	}

	return page.EndPage(systemUsers), nil
}
