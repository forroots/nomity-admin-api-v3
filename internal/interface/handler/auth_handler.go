package handler

import (
	"github.com/forroots/nomity-admin-api-v3/internal/interface/common"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/forroots/nomity-admin-api-v3/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	sessionIdCookieConf common.CookieConfig
	jwtUtil             *jwt.JWTUtil
}

func NewAuthHandler(
	sessionIdCookieConf common.CookieConfig,
	jwtUtil *jwt.JWTUtil,
) *AuthHandler {
	return &AuthHandler{
		sessionIdCookieConf: sessionIdCookieConf,
		jwtUtil:             jwtUtil,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	// リクエストボディをバインド
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorToMiddleware(c, err)
		return
	}

	// ログイン処理を実装
	// ここではダミーのレスポンスを返す
	response.JSON(c, "ログイン成功", nil)
}
