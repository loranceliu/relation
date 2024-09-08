// Code generated by goctl. DO NOT EDIT.
package relation

import (
	"gin/handler/relation"
	"gin/handler/statistics"
	"gin/handler/user"
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRelationRoute(e *gin.Engine) {
	v := e.Group("/v1")
	e.Use(middleware.CorsMiddleware)
	v.Use(middleware.CorsMiddleware,middleware.AuthMiddleware, middleware.DbCheckMiddleware)
	v.POST("/relation", relation.GetRelationListHandle)
	v.POST("/relation/add", relation.AddRelationHandle)
	v.POST("/relation/edit", relation.UpdateRelationHandle)
	v.POST("/relation/delete", relation.DeleteRelationHandle)
	v.POST("/relation/user", relation.GetRelationUserListHandle)
	v.POST("/relation/user/index", relation.GetRelationUserIndexHandle)
	v.POST("/relation/user/add", relation.AddRelationUserHandle)
	v.POST("/relation/user/edit", relation.UpdateRelationUserHandle)
	v.POST("/relation/user/delete", relation.DeleteRelationUserHandle)
	v.POST("/relation/type", relation.GetRelationTypeListHandle)
	v.POST("/relation/type/add", relation.AddRelationTypeHandle)
	v.POST("/relation/type/edit", relation.UpdateRelationTypeHandle)
	v.POST("/relation/type/delete", relation.DeleteRelationTypeHandle)

	v.POST("/user", user.GetSystemUserListHandle)
	v.POST("/user/info", user.GetUserInfoHandle)
	v.POST("/user/mine/info", user.GetUserMineInfoHandle)
	v.POST("/user/add", user.AddUserHandle)
	v.POST("/user/edit", user.UpdateUserHandle)
	v.POST("/user/psersonal/edit", user.UpdatePersonalUserHandle)
	v.POST("/user/password/edit", user.UpdatePasswordUserHandle)
	v.POST("/login", user.LoginHandle)

	v.POST("/statistics/relationTypeNum", statistics.GetRelationTypeNumHandle)
	v.POST("/statistics/relationTypeMoney", statistics.GetRelationTypeMoneyHandle)
	v.POST("/statistics/relationTotalMoney", statistics.GetRelationTotalMoneyHandle)
	v.GET("/statistics/relationUserProfitTop", statistics.GetRelationUserProfitTopHandle)
	v.POST("/statistics/relationCurrentYearTrend", statistics.GetRelationCurrentYearTrendHandle)
	v.POST("/statistics/relationTenYearTrend", statistics.GetRelationTenYearTrendHandle)
}
