package registry

import (
	"database/sql"
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/config"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/db"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail"
)

type Registry struct {
	Config config.Config
	Logger *slog.Logger
	Mailer mail.Mailer
	DB     *sql.DB
}

func NewRegistry(cfg config.Config, logger *slog.Logger) (*Registry, error) {
	mailer := mail.NewSMTPMailer(mail.SMTPMailerConfig{
		Host:       cfg.SMTP.Host,
		Port:       cfg.SMTP.Port,
		Username:   cfg.SMTP.Username,
		Password:   cfg.SMTP.Password,
		From:       cfg.SMTP.From,
		FromName:   cfg.SMTP.FromName,
		DebugPrint: cfg.Mailer.DebugPrint,
		Mock:       cfg.Mailer.Mock,
	})

	db, err := db.NewDB(db.DBParams{
		Driver:   cfg.Database.Driver,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	return &Registry{
		Config: cfg,
		Logger: logger,
		Mailer: mailer,
		DB:     db,
	}, nil
}
