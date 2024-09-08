package response

import (
	"gin/errors"
	"gin/internal/translator"

	"github.com/gin-gonic/gin"
)

// UnifiedResponse 统一返回
type UnifiedResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// HandleResponse 统一返回处理
func HandleResponse(c *gin.Context, data any, err error) {

	if fbError, ok := err.(*errors.ForbiddenError); ok {
		HandleAbortResponse(c, 401, fbError.Error())
		return
	}

	if err != nil {
		c.JSON(200, UnifiedResponse{
			Code:    500,
			Data:    nil,
			Message: translator.Translate(err),
		})
		return
	}

	c.JSON(200, UnifiedResponse{
		Code:    200,
		Data:    data,
		Message: "成功",
	})
}

// HandleAbortResponse 统一 Abort 返回处理
func HandleAbortResponse(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(200, UnifiedResponse{
		Code:    code,
		Data:    nil,
		Message: err,
	})
}

func BusinessAbortResponse(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(code, UnifiedResponse{
		Code:    code,
		Data:    nil,
		Message: err,
	})
}
