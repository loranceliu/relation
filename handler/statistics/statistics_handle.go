package statistics

import (
	"errors"
	"gin/internal/response"
	"gin/logic/statistics"
	"gin/svc"
	"github.com/gin-gonic/gin"
)

func GetRelationTypeNumHandle(c *gin.Context) {
	resp, err := statistics.GetRelationTypeNum(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetRelationTypeMoneyHandle(c *gin.Context) {
	resp, err := statistics.GetRelationTypeMoney(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetRelationTotalMoneyHandle(c *gin.Context) {
	resp, err := statistics.GetRelationTotalMoney(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetRelationCurrentYearTrendHandle(c *gin.Context) {
	resp, err := statistics.GetRelationCurrentYearTrend(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetRelationTenYearTrendHandle(c *gin.Context) {
	resp, err := statistics.GetRelationTenYearTrend(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

func GetRelationUserProfitTopHandle(c *gin.Context) {
	t := c.Query("type")
	if t == "" {
		response.HandleResponse(c, nil, errors.New("参数type不存在"))
		return
	}
	resp, err := statistics.GetRelationUserTopProfit(svc.NewServiceContext(c), t)
	response.HandleResponse(c, resp, err)
}
