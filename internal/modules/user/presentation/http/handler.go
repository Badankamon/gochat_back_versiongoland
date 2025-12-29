package http

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Badankamon/gochat_backend/internal/modules/user/application/dto"
	"github.com/Badankamon/gochat_backend/internal/modules/user/application/usecase"
	"github.com/Badankamon/gochat_backend/internal/shared/errors"
	"github.com/Badankamon/gochat_backend/internal/shared/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	getProfileUC    *usecase.GetProfileUseCase
	updateProfileUC *usecase.UpdateProfileUseCase
	uploadAvatarUC  *usecase.UploadAvatarUseCase
}

func NewUserHandler(
	getProfileUC *usecase.GetProfileUseCase,
	updateProfileUC *usecase.UpdateProfileUseCase,
	uploadAvatarUC *usecase.UploadAvatarUseCase,
) *UserHandler {
	return &UserHandler{
		getProfileUC:    getProfileUC,
		updateProfileUC: updateProfileUC,
		uploadAvatarUC:  uploadAvatarUC,
	}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errors.ErrUnauthorized)
		return
	}

	profile, err := h.getProfileUC.Execute(c.Request.Context(), userID.(string))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Profile retrieved", profile)
}

func (h *UserHandler) UpdateMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errors.ErrUnauthorized)
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.ErrBadRequest)
		return
	}

	profile, err := h.updateProfileUC.Execute(c.Request.Context(), userID.(string), req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Profile updated", profile)
}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errors.ErrUnauthorized)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Error(c, errors.ErrBadRequest)
		return
	}
	defer file.Close()

	// Validate extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		response.Error(c, errors.New(400, "Invalid file type", http.StatusBadRequest))
		return
	}

	url, err := h.uploadAvatarUC.Execute(c.Request.Context(), userID.(string), file, header.Filename)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Avatar uploaded", gin.H{"avatar_url": url})
}
