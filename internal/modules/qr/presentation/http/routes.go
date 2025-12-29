package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *QRHandler, authMiddleware gin.HandlerFunc) {
	qr := router.Group("/qr")
	qr.Use(authMiddleware)
	{
		qr.POST("/user", handler.GenerateUserQR)
		qr.GET("/scan", handler.ScanQR)
	}
}
