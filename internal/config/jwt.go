package config

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey                string `mapstructure:"secret_key"`
	ExpirationMinutes        int    `mapstructure:"expiration_minutes"`
	RefreshExpirationMinutes int    `mapstructure:"refresh_expiration_minutes"`
	SigningAlgorithm         string `mapstructure:"signing_algorithm"`
}

type JWTManager struct {
	SecretKey         string
	Expiration        time.Duration
	RefreshExpiration time.Duration
	SigningAlgorithm  string
}

func NewJWTManager(cfg JWTConfig) *JWTManager {
	manager := &JWTManager{
		SecretKey:         cfg.SecretKey,
		Expiration:        time.Duration(cfg.ExpirationMinutes) * time.Minute,
		RefreshExpiration: time.Duration(cfg.RefreshExpirationMinutes) * time.Minute,
		SigningAlgorithm:  cfg.SigningAlgorithm,
	}

	if manager.SecretKey == "" {
		panic("JWT secret key must be set")
	}
	if manager.Expiration <= 0 {
		panic("JWT expiration must be greater than zero")
	}
	if manager.SigningAlgorithm == "" {
		manager.SigningAlgorithm = "HS256"
	}
	if manager.RefreshExpiration <= 0 {
		manager.RefreshExpiration = manager.Expiration * 2
	}

	return manager
}

func (j *JWTManager) GenerateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID, // int64型で埋め込み
		"exp": time.Now().Add(j.Expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.SigningAlgorithm), claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWTManager) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(j.RefreshExpiration).Unix(),
		"typ": "refresh",
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.SigningAlgorithm), claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWTManager) ParseToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid subject type")
	}

	return int64(subFloat), nil
}

func (j *JWTManager) ParseRefreshToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	typ, ok := claims["typ"].(string)
	if !ok || typ != "refresh" {
		return 0, errors.New("invalid token type")
	}

	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	return int64(subFloat), nil
}
