package config

type SendGridConfig struct {
	APIKey    string `mapstructure:"api_key"`
	EmailFrom string `mapstructure:"email_from"`
	NameFrom  string `mapstructure:"name_from"`

	DebugPrint bool `mapstructure:"debug_print"`
	Mock       bool `mapstructure:"mock"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
	FromName string `mapstructure:"from_name"`

	DebugPrint bool `mapstructure:"debug_print"`
	Mock       bool `mapstructure:"mock"`
}
