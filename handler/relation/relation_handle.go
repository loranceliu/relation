package relation

import (
	"errors"
	"gin/internal/response"
	relationLogic "gin/logic/relation"
	"gin/svc"
	"gin/types/relation"

	"github.com/gin-gonic/gin"
)

// AddRelationHandle 添加人情
func AddRelationHandle(c *gin.Context) {
	var req relation.RelationRequest
	var err error

	err = c.ShouldBind(&req)

	if req.RelationUserID == 0 {
		err = errors.New("请选择账单对象")
	}

	if req.RelationTypeID == 0 {
		err = errors.New("请选择账单类型")
	}

	if req.TransactionType == 0 {
		err = errors.New("请选择账单类目")
	}

	if req.Money == 0 {
		err = errors.New("请填写账单金额")
	}

	if req.Date == "" {
		err = errors.New("请选择账单日期")
	}

	if err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := relationLogic.AddRelation(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// UpdateRelationHandle 更新人情
func UpdateRelationHandle(c *gin.Context) {
	var req relation.RelationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.UpdateRelation(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// DeleteRelationHandle 删除人情
func DeleteRelationHandle(c *gin.Context) {
	var req relation.DeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.DeleteRelation(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// GetRelationListHandle 获取人情列表
func GetRelationListHandle(c *gin.Context) {
	var req relation.RelationPageRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.GetRelationList(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// GetRelationUserListHandle 获取人情用户列表
func GetRelationUserListHandle(c *gin.Context) {
	var req relation.RelationUserPageRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.GetRelationUserList(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func GetRelationUserIndexHandle(c *gin.Context) {
	var req relation.RelationUserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.GetRelationUserIndex(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// AddRelationUserHandle 添加人情用户
func AddRelationUserHandle(c *gin.Context) {
	var req relation.RelationUserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.AddRelationUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func DeleteRelationUserHandle(c *gin.Context) {
	var req relation.DeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.DeleteRelationUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// UpdateRelationUserHandle 更新人情用户
func UpdateRelationUserHandle(c *gin.Context) {
	var req relation.RelationUserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.UpdateRelationUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func AddRelationTypeHandle(c *gin.Context) {
	var req relation.RelationTypeRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.AddRelationType(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func UpdateRelationTypeHandle(c *gin.Context) {
	var req relation.RelationTypeRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.UpdateRelationType(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func DeleteRelationTypeHandle(c *gin.Context) {
	var req relation.DeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := relationLogic.DeleteRelationType(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

// GetRelationTypeListHandle 获取人情类型
func GetRelationTypeListHandle(c *gin.Context) {
	resp, err := relationLogic.GetRelationTypeList(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
