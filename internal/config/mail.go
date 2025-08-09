package config

type MailerConfig struct {
	DebugPrint       bool     `mapstructure:"debug_print"`
	Mock             bool     `mapstructure:"mock"`
	NotifyRecipients []string `mapstructure:"notify_recipients"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
	FromName string `mapstructure:"from_name"`
}
