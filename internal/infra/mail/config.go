package mail

type SMTPMailerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string

	DebugPrint bool // ログ出力するかどうか
	Mock       bool // モックモードかどうか
}
