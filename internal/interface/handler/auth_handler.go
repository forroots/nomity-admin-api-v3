package handler

import "github.com/gin-gonic/gin"

type AuthHandler struct {
	// 依存関係を追加
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		// 依存関係を追加
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// ログイン処理を実装
	// ここではダミーのレスポンスを返す
	c.JSON(200, gin.H{"message": "Login successful", "email": req.Email})
}
