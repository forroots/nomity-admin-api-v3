package db

import "fmt"

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	Debug    bool
}

func (p DBConfig) DSN() string {
	switch p.Driver {
	case "postgres":
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			p.User, p.Password, p.Host, p.Port, p.DBName, p.SSLMode,
		)
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			p.User, p.Password, p.Host, p.Port, p.DBName,
		)
	default:
		return ""
	}
}
