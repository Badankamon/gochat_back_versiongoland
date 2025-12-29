package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/verification-code", handler.SendVerificationCode)
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
