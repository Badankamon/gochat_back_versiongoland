package http

import (
	"github.com/Badankamon/gochat_backend/internal/modules/auth/application/dto"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/application/usecase"
	"github.com/Badankamon/gochat_backend/internal/shared/errors"
	"github.com/Badankamon/gochat_backend/internal/shared/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	sendCodeUC *usecase.SendVerificationCodeUseCase
	registerUC *usecase.RegisterUseCase
	loginUC    *usecase.LoginUseCase
}

func NewAuthHandler(
	sendCodeUC *usecase.SendVerificationCodeUseCase,
	registerUC *usecase.RegisterUseCase,
	loginUC *usecase.LoginUseCase,
) *AuthHandler {
	return &AuthHandler{
		sendCodeUC: sendCodeUC,
		registerUC: registerUC,
		loginUC:    loginUC,
	}
}

func (h *AuthHandler) SendVerificationCode(c *gin.Context) {
	var req dto.SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.ErrBadRequest)
		return
	}

	if err := h.sendCodeUC.Execute(c.Request.Context(), req.Phone); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Verification code sent", nil)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.ErrBadRequest)
		return
	}

	resp, err := h.registerUC.Execute(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "User registered successfully", resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.ErrBadRequest)
		return
	}

	resp, err := h.loginUC.Execute(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Login successful", resp)
}
