package handler

import (
	"blog/internal/dto"
	"blog/internal/service"
	"blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "请输入用户名和密码")
		return
	}

	resp, err := h.authService.Login(req)
	if err != nil {
		utils.Error(c, 401, err.Error())
		return
	}

	utils.Success(c, resp)
}

func (h *AuthHandler) Profile(c *gin.Context) {
	userID := c.GetUint("user_id")
	resp, err := h.authService.GetProfile(userID)
	if err != nil {
		utils.Error(c, 404, err.Error())
		return
	}
	utils.Success(c, resp)
}
