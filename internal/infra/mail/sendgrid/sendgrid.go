package sendgrid

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailerConfig struct {
	APIKey    string
	FromEmail string
	FromName  string

	DebugPrint bool // ログ出力するかどうか
	Mock       bool // モックモードかどうか
}

type SendGridMailer struct {
	client     *sendgrid.Client
	fromEmail  *mail.Email
	DebugPrint bool // ログ出力するかどうか
	Mock       bool // モックモードかどうか
}

func NewSendGridSender(cfg SendGridMailerConfig) *SendGridMailer {
	client := sendgrid.NewSendClient(cfg.APIKey)
	fromEmail := mail.NewEmail(cfg.FromName, cfg.FromEmail)
	return &SendGridMailer{
		client:     client,
		fromEmail:  fromEmail,
		DebugPrint: cfg.DebugPrint,
		Mock:       cfg.Mock,
	}
}

func (s *SendGridMailer) SendSingleEmail(to string, subject, plainTextContent, htmlContent string) (*rest.Response, error) {
	toEmail := mail.NewEmail("", to)

	message := mail.NewSingleEmail(s.fromEmail, subject, toEmail, plainTextContent, htmlContent)
	res, err := s.client.Send(message)
	return res, err
}
