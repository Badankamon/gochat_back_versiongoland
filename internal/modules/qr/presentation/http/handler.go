package http

import (
	"github.com/Badankamon/gochat_backend/internal/modules/qr/application/usecase"
	"github.com/Badankamon/gochat_backend/internal/shared/errors"
	"github.com/Badankamon/gochat_backend/internal/shared/response"
	"github.com/Badankamon/gochat_backend/pkg/qrcode"
	"github.com/gin-gonic/gin"
)

type QRHandler struct {
	generateUC *usecase.GenerateTicketUseCase
	scanUC     *usecase.ScanTicketUseCase
}

func NewQRHandler(genUC *usecase.GenerateTicketUseCase, scanUC *usecase.ScanTicketUseCase) *QRHandler {
	return &QRHandler{
		generateUC: genUC,
		scanUC:     scanUC,
	}
}

func (h *QRHandler) GenerateUserQR(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errors.ErrUnauthorized)
		return
	}

	// For User QR, target is self. Mode defaults to permanent or user choice.
	// Implementing MVP: Permanent User QR
	ticket, err := h.generateUC.Execute(c.Request.Context(), userID.(string), "user", userID.(string), "permanent")
	if err != nil {
		response.Error(c, err)
		return
	}

	// Generate real image URL
	qrURL := "https://api.gochat.com/qr/scan?ticket=" + ticket.Ticket

	// Generate QR Image (Base64)
	qrBase64, err := qrcode.GenerateBase64(qrURL, 256)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "QR Ticket generated", gin.H{
		"ticket":     ticket.Ticket,
		"qr_url":     qrURL,
		"qr_image":   "data:image/png;base64," + qrBase64,
		"expires_at": ticket.ExpiresAt,
	})
}

func (h *QRHandler) ScanQR(c *gin.Context) {
	ticketStr := c.Query("ticket")
	if ticketStr == "" {
		response.Error(c, errors.New(400, "ticket is required", 400))
		return
	}

	result, err := h.scanUC.Execute(c.Request.Context(), ticketStr)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "QR Scanned successfully", result)
}
