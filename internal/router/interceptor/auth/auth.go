package auth

import (
	"api/configs"
	"api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := &Auth{}
		err := c.BindHeader(auth)
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		if auth.Username != configs.Get().Auth.Username || auth.Password != configs.Get().Auth.Password {
			response.Fail(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
