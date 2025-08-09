package registry

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/forroots/nomity-admin-api-v3/internal/config"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/db"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail/sendgrid"
	"github.com/forroots/nomity-admin-api-v3/internal/infra/mail/smtp"
	"github.com/forroots/nomity-admin-api-v3/internal/utils/jwt"
)

type Registry struct {
	Config  config.Config
	Logger  *slog.Logger
	Mailer  mail.IMailer
	DB      *sql.DB
	JWTUtil *jwt.JWTUtil
}

func NewRegistry(cfg config.Config, logger *slog.Logger) (*Registry, error) {
	mailer := mail.NewMailer(
		sendgrid.SendGridMailerConfig{},
		smtp.SMTPMailerConfig{},
		logger,
	)

	db, err := db.NewDB(db.DBConfig{
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

	// JWTのセットアップ
	jwtConfig := jwt.Config{
		SecretKey:         cfg.JWT.SecretKey,
		Expiration:        time.Duration(cfg.JWT.ExpirationMinutes) * time.Minute,
		RefreshExpiration: time.Duration(cfg.JWT.RefreshExpirationMinutes) * time.Minute,
		SigningAlgorithm:  cfg.JWT.SigningAlgorithm,
	}
	jwtUtil := jwt.NewJWTUtil(jwtConfig)

	return &Registry{
		Config:  cfg,
		Logger:  logger,
		Mailer:  mailer,
		DB:      db,
		JWTUtil: jwtUtil,
	}, nil
}
