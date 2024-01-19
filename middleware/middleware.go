package middleware

import (
	"TodoList/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// token 验证 Json Web Token
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404 // 找不到 token
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 假token，无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 // Token 超时无效
			}
		}

		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "token 解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
