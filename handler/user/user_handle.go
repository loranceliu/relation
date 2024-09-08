package user

import (
	"gin/internal/response"
	userLogic "gin/logic/user"
	"gin/svc"
	"gin/types/relation"
	user "gin/types/user"
	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {
	var req user.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.Login(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func GetUserInfoHandle(c *gin.Context) {
	resp, err := userLogic.GetUserInfo(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetUserMineInfoHandle(c *gin.Context) {
	resp, err := userLogic.GetUserMineInfo(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func AddUserHandle(c *gin.Context) {
	var req user.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.AddUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func UpdateUserHandle(c *gin.Context) {
	var req user.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.UpdateUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func UpdatePersonalUserHandle(c *gin.Context) {
	var req user.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.UpdatePersonalUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func UpdatePasswordUserHandle(c *gin.Context) {
	var req user.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.UpdatePasswordUser(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}

func GetSystemUserListHandle(c *gin.Context) {
	var req relation.SystemUserPageRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := userLogic.GetSystemUserList(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
