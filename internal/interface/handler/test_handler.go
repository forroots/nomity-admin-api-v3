package handler

import (
	"fmt"
	"net/http"

	"github.com/forroots/nomity-admin-api-v3/internal/application"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	Mailer mail.Mailer
}

func NewMailHandler(mailer mail.Mailer) *TestHandler {
	return &TestHandler{
		Mailer: mailer,
	}
}

func (h *TestHandler) PanicTest(c *gin.Context) {
	// 意図的にパニックを発生させる
	panic("This is a test panic")
}

func (h *TestHandler) ErrorTest(c *gin.Context) {
	// 意図的にエラーを発生させる
	err := fmt.Errorf("this is a test errorよ")
	c.Error(err) // エラーをコンテキストに追加
}

func (h *TestHandler) AppErrorTest(c *gin.Context) {
	// 意図的にエラーを発生させる
	err := fmt.Errorf("this is a test app errorよ")
	appErr := application.NewAppError(
		application.WithHttpStatus(http.StatusBadRequest),
		application.WithCode("test_error"),
		application.WithMessage("This is a test app error"),
		application.WithError(err),
	)
	c.Error(appErr) // エラーをコンテキストに追加
}

func (h *TestHandler) SendTestEmail(c *gin.Context) {
	err := h.Mailer.Send(
		[]string{"kcntsurk@gmail.com"}, // ← 実際の宛先に置き換えてください
		nil,
		"テストメール",
		"これはテストメールですよ。",
	)
	if err != nil {
		appErr := application.NewAppError(
			application.WithHttpStatus(http.StatusInternalServerError),
			application.WithCode("mail_send_error"),
			application.WithMessage("メール送信に失敗しました"),
			application.WithError(err),
		)
		c.Error(appErr) // エラーをコンテキストに追加
		return
	}

	response.JSON(c, "メール送信に成功しました", nil)
}

func (h *TestHandler) JustPost(c *gin.Context) {
	// POSTリクエストのテスト用ハンドラ
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		appErr := application.NewAppError(
			application.WithHttpStatus(http.StatusBadRequest),
			application.WithCode("invalid_request_body"),
			application.WithMessage("無効なリクエストボディ"),
			application.WithError(err),
		)
		c.Error(appErr) // エラーをコンテキストに追加
		return
	}

	response.JSON(c, "POSTリクエストを受け取りました", requestBody)
}
