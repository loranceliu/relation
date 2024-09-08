package middleware

import (
	"github.com/gin-gonic/gin"
)

func DbCheckMiddleware(c *gin.Context) {

	c.Next()
}
