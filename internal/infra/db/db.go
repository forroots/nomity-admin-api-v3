package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQLドライバ
)

func NewDB(p DBConfig) (*sql.DB, error) {
	var dsn string
	switch p.Driver {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode,
		)
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			p.User, p.Password, p.Host, p.Port, p.DBName,
		)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", p.Driver)
	}

	db, err := sql.Open(p.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}
