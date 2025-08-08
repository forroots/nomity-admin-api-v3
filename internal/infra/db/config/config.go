// internal/config/config.go
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env string `mapstructure:"env"`
	} `mapstructure:"app"`

	Server struct {
		Port           int      `mapstructure:"port"`
		BasePath       string   `mapstructure:"base_path"`
		TrustedProxies []string `mapstructure:"trusted_proxies"`
	} `mapstructure:"server"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
		Debug    bool   `mapstructure:"debug"`
	} `mapstructure:"database"`

	Cookie struct {
		SessionID CookieConfig `mapstructure:"session_id"`
		CSRFToken CookieConfig `mapstructure:"csrf_token"`
	} `mapstructure:"cookie"`

	CORS struct {
		Enabled              bool     `mapstructure:"enabled"`
		AllowOrigins         []string `mapstructure:"allow_origins"`
		AdvancedAllowHeaders []string `mapstructure:"advanced_allow_headers"`
	} `mapstructure:"cors"`

	CSRF struct {
		Enabled         bool   `mapstructure:"enabled"`
		TokenCookieName string `mapstructure:"token_cookie_name"`
		TokenHeader     string `mapstructure:"token_header"`
	} `mapstructure:"csrf"`

	Mailer struct {
		DebugPrint       bool     `mapstructure:"debug_print"`
		Mock             bool     `mapstructure:"mock"`
		NotifyRecipients []string `mapstructure:"notify_recipients"`
	} `mapstructure:"mailer"`

	SMTP struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		From     string `mapstructure:"from"`
		FromName string `mapstructure:"from_name"`
	} `mapstructure:"smtp"`

	Google struct {
		CredentialsPath string `mapstructure:"credentials_path"`
		SpreadsheetID   string `mapstructure:"spreadsheet_id"`
		SpreadsheetName string `mapstructure:"spreadsheet_name"`
	} `mapstructure:"google"`
}

type CookieConfig struct {
	Name     string `mapstructure:"name"`
	Path     string `mapstructure:"path"`
	Domain   string `mapstructure:"domain"`
	Secure   bool   `mapstructure:"secure"`
	HttpOnly bool   `mapstructure:"http_only"`
	SameSite string `mapstructure:"same_site"`
	MaxAge   int    `mapstructure:"max_age"`
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) EnvDevelopment() bool {
	return c.App.Env == "development"
}
func (c *Config) EnvProduction() bool {
	return c.App.Env == "production"
}
func (c *Config) EnvTest() bool {
	return c.App.Env == "test"
}
