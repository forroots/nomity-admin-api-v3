package config

import "github.com/forroots/nomity-admin-api-v3/internal/infra/db"

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	Debug    bool   `mapstructure:"debug"`
}

func (c *DatabaseConfig) ToInfraDBConfig() db.DBConfig {
	return db.DBConfig{
		Driver:   c.Driver,
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		DBName:   c.DBName,
		SSLMode:  c.SSLMode,
		Debug:    c.Debug,
	}
}
