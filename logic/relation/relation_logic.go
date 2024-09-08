package relation

import (
	"errors"
	"fmt"
	"gin/entity/vo"
	"gin/model"
	"gin/svc"
	"gin/types/relation"
	"gin/utils"
	"github.com/jinzhu/copier"
	"github.com/mozillazg/go-pinyin"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func GetRelationList(ctx *svc.ServiceContext, req *relation.RelationPageRequest) (*vo.PageVO, error) {

	// 构建原生 SQL 查询语句
	sqlQuery := "SELECT tr.relation_id, tr.relation_name,tr.relation_type_id,tr.relation_user_id,tr.money, tr.remark,trt.relation_type_name, tru.relation_user_name, tr.transaction_type, tr.date FROM tb_relation tr LEFT JOIN tb_relation_type trt ON tr.relation_type_id = trt.relation_type_id LEFT JOIN tb_relation_user tru ON tr.relation_user_id = tru.relation_user_id where tru.status = 1"
	var sqlArgs []interface{}

	userId := ctx.Value("user_id")
	sqlQuery += " AND tr.owner_id = " + userId

	// 添加过滤条件
	if req.RelationUserId != 0 {
		sqlQuery += " AND tr.relation_user_id = ?"
		sqlArgs = append(sqlArgs, req.RelationUserId)
	}

	if req.TransactionType != 0 {
		sqlQuery += " AND tr.transaction_type = ?"
		sqlArgs = append(sqlArgs, req.TransactionType)
	}

	if req.RelationTypeId != 0 {
		sqlQuery += " AND trt.relation_type_id = ?"
		sqlArgs = append(sqlArgs, req.RelationTypeId)
	}

	if req.StartTime != "" && req.EndTime != "" {
		sqlQuery += " AND tr.date >= ? And tr.date <= ?"
		sqlArgs = append(sqlArgs, req.StartTime, req.EndTime)
	}

	if req.Remark != "" {
		sqlQuery += fmt.Sprintf(" AND tr.remark like '%%%s%%'", req.Remark)
	}

	sqlQuery += " Order By tr.relation_id desc"

	page := utils.New(sqlQuery, sqlArgs, &req.PageRequest)
	sql := page.StartPage()
	var relations []vo.RelationVO
	if err := model.DB.Raw(sql, sqlArgs...).Scan(&relations).Error; err != nil {
		return nil, err
	}
	return page.EndPage(relations), nil
}

func GetRelationUserList(ctx *svc.ServiceContext, req *relation.RelationUserPageRequest) (*vo.PageVO, error) {
	// 构建原生 SQL 查询语句

	userId := ctx.Value("user_id")

	var builder strings.Builder

	_, _ = fmt.Fprintf(&builder, "SELECT tru.relation_user_id, tru.status, tru.relation_user_name,tru.remark, tru.sex,IFNULL((SELECT SUM(money) FROM tb_relation tr WHERE tr.relation_user_id = tru.relation_user_id AND tr.transaction_type = 1 AND owner_id = %s), 0) AS income,IFNULL((SELECT SUM(money) FROM tb_relation tr WHERE tr.relation_user_id = tru.relation_user_id AND tr.transaction_type = 2 and owner_id = %s), 0) AS expend FROM tb_relation_user tru where tru.owner_id = %s", userId, userId, userId)
	sqlQuery := builder.String()
	var sqlArgs []interface{}

	// 添加过滤条件
	if req.RelationUserId != 0 {
		sqlQuery += " AND relation_user_id = ?"
		sqlArgs = append(sqlArgs, req.RelationUserId)
	}

	if req.Status != 0 {
		sqlQuery += " AND status = ?"
		sqlArgs = append(sqlArgs, req.Status)
	}

	sqlQuery += " Order By expend desc"

	page := utils.New(sqlQuery, sqlArgs, &req.PageRequest)
	sql := page.StartPage()

	var relations []vo.RelationUserVO
	if err := model.DB.Raw(sql, sqlArgs...).Scan(&relations).Error; err != nil {
		return nil, err
	}

	return page.EndPage(relations), nil
}

func GetRelationUserIndex(ctx *svc.ServiceContext, req *relation.RelationUserRequest) (*vo.RelationUserIndexVO, error) {
	// 构建原生 SQL 查询语句

	userId := ctx.Value("user_id")

	var builder strings.Builder

	_, _ = fmt.Fprintf(&builder, "SELECT tru.relation_user_id, tru.status, tru.relation_user_name,tru.remark, tru.prefix, tru.sex,IFNULL((SELECT SUM(money) FROM tb_relation tr WHERE tr.relation_user_id = tru.relation_user_id AND tr.transaction_type = 1 AND owner_id = %s), 0) AS income,IFNULL((SELECT SUM(money) FROM tb_relation tr WHERE tr.relation_user_id = tru.relation_user_id AND tr.transaction_type = 2 and owner_id = %s), 0) AS expend FROM tb_relation_user tru where tru.owner_id = %s", userId, userId, userId)
	sqlQuery := builder.String()

	if req.Search != "" {
		sqlQuery += fmt.Sprintf(" AND tru.relation_user_name like '%%%s%%'", req.Search)
	}

	var relations []vo.RelationUserVO
	if err := model.DB.Raw(sqlQuery).Scan(&relations).Error; err != nil {
		return nil, err
	}

	var item []vo.RelationUserItemVO
	var index []string

	indexData := vo.RelationUserIndexVO{
		Index: index,
		Item:  item,
	}

	if len(relations) > 0 {

		for _, item := range relations {
			found := false
			for _, i := range index {
				if item.Prefix == i {
					found = true
					break
				}
			}

			if !found {
				index = append(index, item.Prefix)
			}
		}

		sort.Slice(index, func(i, j int) bool {
			return index[i] < index[j]
		})

		for _, s := range index {

			var v vo.RelationUserItemVO
			v.Data = []vo.RelationUserVO{}

			for _, item := range relations {
				if s == item.Prefix {
					v.Type = s
					v.Data = append(v.Data, item)
				}
			}

			item = append(item, v)
		}

		indexData.Index = index
		indexData.Item = item
	}

	return &indexData, nil
}

func AddRelation(ctx *svc.ServiceContext, req *relation.RelationRequest) (inter interface{}, err error) {
	rl := model.Relation{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	userId := ctx.Value("user_id")
	rl.OwnerID, _ = strconv.Atoi(userId)
	model.DB.Create(&rl)
	return
}

func UpdateRelation(ctx *svc.ServiceContext, req *relation.RelationRequest) (inter interface{}, err error) {
	rl := model.Relation{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	model.DB.Updates(&rl)
	return
}

func DeleteRelation(ctx *svc.ServiceContext, req *relation.DeleteRequest) (inter interface{}, err error) {
	rl := model.Relation{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	model.DB.Where("relation_id in ?", req.Ids).Delete(&rl)
	return
}

func DeleteRelationType(ctx *svc.ServiceContext, req *relation.DeleteRequest) (inter interface{}, err error) {

	var rst int

	if err := model.DB.Raw("select count(*) from tb_relation where relation_type_id in ?", req.Ids).Scan(&rst).Error; err != nil {
		return nil, err
	}

	if rst > 0 {
		return nil, errors.New("此类目已使用，请解除后再删除")
	}

	rl := model.RelationType{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	model.DB.Where("relation_type_id in ?", req.Ids).Delete(&rl)
	return
}

func AddRelationUser(ctx *svc.ServiceContext, req *relation.RelationUserRequest) (inter interface{}, err error) {
	rl := model.RelationUser{}
	err = copier.Copy(&rl, req)

	rl.Prefix = getFirstLetter(rl.RelationUserName)

	if err != nil {
		return nil, err
	}
	userId := ctx.Value("user_id")
	rl.OwnerID, _ = strconv.Atoi(userId)
	model.DB.Create(&rl)
	return
}

func DeleteRelationUser(ctx *svc.ServiceContext, req *relation.DeleteRequest) (inter interface{}, err error) {

	var rst int

	if err := model.DB.Raw("select count(*) from tb_relation where relation_user_id in ?", req.Ids).Scan(&rst).Error; err != nil {
		return nil, err
	}

	if rst > 0 {
		return nil, errors.New("此对象已使用，请解除后再删除")
	}

	rl := model.RelationUser{}
	err = copier.Copy(&rl, req)
	if err != nil {
		return nil, err
	}
	model.DB.Where("relation_user_id in ?", req.Ids).Delete(&rl)
	return
}

func AddRelationType(ctx *svc.ServiceContext, req *relation.RelationTypeRequest) (inter interface{}, err error) {
	rl := model.RelationType{}
	err = copier.Copy(&rl, req)

	if err != nil {
		return nil, err
	}
	userId := ctx.Value("user_id")
	rl.OwnerID, _ = strconv.Atoi(userId)
	model.DB.Create(&rl)
	return
}

func UpdateRelationType(ctx *svc.ServiceContext, req *relation.RelationTypeRequest) (inter interface{}, err error) {
	rl := model.RelationType{}
	err = copier.Copy(&rl, req)

	if err != nil {
		return nil, err
	}
	model.DB.Updates(&rl)
	return
}

func getFirstLetter(str string) string {
	firstRune, _ := utf8.DecodeRuneInString(str)
	if unicode.Is(unicode.Han, firstRune) {

		p := pinyin.NewArgs()
		result := pinyin.LazyPinyin(string(firstRune), p)

		firstLetters := ""
		for _, py := range result {
			firstLetters += string(py[0])
		}

		return strings.ToUpper(string(firstLetters[0]))
	} else {
		return strings.ToUpper(string(firstRune))
	}
}

func UpdateRelationUser(ctx *svc.ServiceContext, req *relation.RelationUserRequest) (inter interface{}, err error) {
	rl := model.RelationUser{}
	err = copier.Copy(&rl, req)

	rl.Prefix = getFirstLetter(rl.RelationUserName)

	if err != nil {
		return nil, err
	}
	model.DB.Updates(&rl)
	return
}

func GetRelationTypeList(ctx *svc.ServiceContext) (rlVo *[]vo.RelationTypeVO, err error) {
	userId := ctx.Value("user_id")

	var builder strings.Builder

	_, _ = fmt.Fprintf(&builder, "select relation_type_id,relation_type_name from tb_relation_type where owner_id = %s", userId)
	sqlQuery := builder.String()

	var relations []vo.RelationTypeVO
	if err := model.DB.Raw(sqlQuery).Scan(&relations).Error; err != nil {
		return nil, err
	}
	return &relations, nil
}
