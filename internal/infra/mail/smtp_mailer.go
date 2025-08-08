package mail

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"strings"
)

const base64LineLength = 76 // Base64での1行の最大文字数（RFC準拠）

type SMTPMailer struct {
	config SMTPMailerConfig
}

func NewSMTPMailer(cfg SMTPMailerConfig) *SMTPMailer {
	return &SMTPMailer{
		config: cfg,
	}
}

func (m *SMTPMailer) Send(to []string, cc []string, subject string, body string) error {
	if m.config.DebugPrint {
		fmt.Printf("Sending email:\nTo: %s\nCc: %s\nSubject: %s\nBody: %s\n",
			strings.Join(to, ", "), strings.Join(cc, ", "), subject, body)
	}

	addr := fmt.Sprintf("%s:%d", m.config.Host, m.config.Port)
	auth := smtp.PlainAuth("", m.config.Username, m.config.Password, m.config.Host)

	// 差出人（名前付き）
	from := m.config.From
	fullFrom := from
	if m.config.FromName != "" {
		encodedFromName := mime.BEncoding.Encode("utf-8", m.config.FromName)
		fullFrom = fmt.Sprintf("%s <%s>", encodedFromName, from)
	}

	// 件名エンコード
	encodedSubject := mime.BEncoding.Encode("utf-8", subject)

	// ヘッダ定義
	headers := map[string]string{
		"From":                      fullFrom,
		"To":                        strings.Join(to, ", "),
		"Subject":                   encodedSubject,
		"MIME-Version":              "1.0",
		"Content-Type":              `text/plain; charset="UTF-8"`,
		"Content-Transfer-Encoding": "base64",
	}
	if len(cc) > 0 {
		headers["Cc"] = strings.Join(cc, ", ")
	}

	// メール本文をbase64でエンコードし、76文字ごとに改行
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(wrapBase64(body))

	// 宛先（To + Cc）
	recipients := append([]string{}, to...)
	recipients = append(recipients, cc...)

	if m.config.Mock {
		// モックモードでは実際の送信は行わない
		fmt.Println("Mock mode enabled, email not sent.")
		return nil
	}

	return smtp.SendMail(addr, auth, from, recipients, []byte(msg.String()))
}

// 本文をbase64エンコードして76文字ごとに折り返し
func wrapBase64(body string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(body))
	var b strings.Builder
	for i := 0; i < len(encoded); i += base64LineLength {
		end := i + base64LineLength
		if end > len(encoded) {
			end = len(encoded)
		}
		b.WriteString(encoded[i:end] + "\r\n")
	}
	return b.String()
}
