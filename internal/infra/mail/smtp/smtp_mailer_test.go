package smtp

import (
	"testing"
)

func TestSMTPMailer_Send_Real(t *testing.T) {
	cfg := SMTPMailerConfig{
		Host:       "smtp.gmail.com",
		Port:       587,
		Username:   "info@diverlabo-support.com",
		Password:   "fripnntdhdpjezaw",
		From:       "info@diverlabo-support.com",
		FromName:   "Nomity（dev）",
		DebugPrint: true, // ログ出力を有効にする
		Mock:       true, // モックモードを有効にする
	}

	mailer := NewSMTPMailer(cfg) // ログ出力を有効にし、モックモードは無効

	err := mailer.SendSingleEmail(
		"kcntsurk@gmail.com",
		"integration test",
		"これはテストメールです！",
		"<p>これはテストメールです！</p>",
	)

	if err != nil {
		t.Fatalf("メール送信に失敗しました: %v", err)
	}
}
