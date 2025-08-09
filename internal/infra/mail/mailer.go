package mail

import (
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail/sendgrid"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail/smtp"
)

type IMailer interface {
	SendSingleEmail(to string, subject, plainTextContent, htmlContent string) error
}

type Mailer struct {
	sendGridMailer *sendgrid.SendGridMailer
	smtpMailer     *smtp.SMTPMailer
	logger         *slog.Logger
}

func NewMailer(sendGridConfig sendgrid.SendGridMailerConfig, smtpConfig smtp.SMTPMailerConfig, logger *slog.Logger) IMailer {
	return &Mailer{
		sendGridMailer: sendgrid.NewSendGridSender(sendGridConfig),
		smtpMailer:     smtp.NewSMTPMailer(smtpConfig),
		logger:         logger,
	}
}

func (s *Mailer) SendSingleEmail(to string, subject string, plainTextContent string, htmlContent string) error {
	if !checkDomainForSendGrid(to) {
		// SendGridで送信できないドメインのため、SMTPから送信します
		err := s.smtpMailer.Send([]string{to}, nil, subject, plainTextContent)
		return err
	} else {
		// SendGridで送信します
		_, err := s.sendGridMailer.SendSingleEmail(to, subject, plainTextContent, htmlContent)
		if err != nil {
			// SendGridから送信失敗した場合はSMTPで再送信する
			err = s.smtpMailer.Send([]string{to}, nil, subject, plainTextContent)
		}
		return err
	}
}

// checkDomainForSendGrid : SendGridで送信できないドメインであればfalseを返す
func checkDomainForSendGrid(to string) bool {
	sendgridCantSendDomains := []string{"ezweb.ne.jp", "au.com", "hotmail.co.jp"}
	for _, cantSendDomain := range sendgridCantSendDomains {
		// toの@以降がcantSendDomainと一致する場合はNG
		if len(to) > len(cantSendDomain) && to[len(to)-len(cantSendDomain):] == cantSendDomain {
			return false
		}
	}
	return true
}
