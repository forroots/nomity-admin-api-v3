package application

type Mailer interface {
	Send(to []string, cc []string, subject string, body string) error
}
