package middleware

import (
	"gin/internal/response"
	"gin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {

	requestPath := c.Request.URL.Path

	if isExcludedURL(requestPath) {
		c.Next()
		return
	}

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		response.HandleAbortResponse(c, 403, "拒绝访问")
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 根据你的需求，返回用于验证签名的密钥
		return []byte(utils.Secret), nil
	})

	if err != nil || !token.Valid {
		response.HandleAbortResponse(c, 403, "拒绝访问")
		c.Abort()
		return
	}

	// 在请求上下文中添加用户信息
	claims, _ := token.Claims.(jwt.MapClaims)
	c.Set("user_id", claims["user_id"])

	c.Next()
}

func isExcludedURL(url string) bool {
	// 在这里添加你要排除的 URL，例如登录、注册等
	excludedURLs := []string{"/v1/login", "/v1/register"}

	for _, excluded := range excludedURLs {
		if url == excluded {
			return true
		}
	}

	return false
}
