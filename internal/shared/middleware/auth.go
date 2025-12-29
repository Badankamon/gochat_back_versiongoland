package middleware

import (
	"net/http"
	"strings"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/Badankamon/gochat_backend/internal/shared/crypto"
	"github.com/Badankamon/gochat_backend/internal/shared/errors"
	"github.com/Badankamon/gochat_backend/internal/shared/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, errors.New(1004, "Authorization header required", http.StatusUnauthorized))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, errors.New(1004, "Invalid authorization format", http.StatusUnauthorized))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := crypto.ValidateToken(tokenString, cfg)
		if err != nil {
			response.Error(c, errors.New(1004, "Invalid or expired token", http.StatusUnauthorized))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
