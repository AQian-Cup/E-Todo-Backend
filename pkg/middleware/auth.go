package middleware

import (
	"crypto/ecdsa"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/jwt"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(key *ecdsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.GetHeader("Authorization")
		name, err := jwt.Validate(s, key)
		if err != nil {
			response.Write(c, errno.UserAuthError)
			c.Abort()
			return
		}
		c.Set("name", name)
		println("123")
		c.Next()
	}
}
