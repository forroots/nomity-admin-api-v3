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

	Server ServerConfig `mapstructure:"server"`
	Cookie struct {
		SessionID CookieConfig `mapstructure:"session_id"`
		CSRFToken CookieConfig `mapstructure:"csrf_token"`
	} `mapstructure:"cookie"`
	CORS CORSConfig `mapstructure:"cors"`
	CSRF CSRFConfig `mapstructure:"csrf"`

	JWT JWTConfig `mapstructure:"jwt"`

	Database DatabaseConfig `mapstructure:"database"`

	SendGridConfig SendGridConfig `mapstructure:"sendgrid"`
	SMTPConfig     SMTPConfig     `mapstructure:"smtp"`
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
