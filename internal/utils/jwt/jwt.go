package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SecretKey         string
	Expiration        time.Duration
	RefreshExpiration time.Duration
	SigningAlgorithm  string
}

type JWTUtil struct {
	config Config
}

func NewJWTUtil(cfg Config) *JWTUtil {
	if cfg.SecretKey == "" {
		panic("JWT secret key must be set")
	}
	if cfg.Expiration <= 0 {
		panic("JWT expiration must be greater than zero")
	}
	if cfg.SigningAlgorithm == "" {
		cfg.SigningAlgorithm = "HS256"
	}
	if cfg.RefreshExpiration <= 0 {
		cfg.RefreshExpiration = cfg.Expiration * 2
	}
	return &JWTUtil{config: cfg}
}

func (j *JWTUtil) GenerateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID, // int64型で埋め込み
		"exp": time.Now().Add(j.config.Expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.config.SigningAlgorithm), claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

func (j *JWTUtil) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(j.config.RefreshExpiration).Unix(),
		"typ": "refresh",
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.config.SigningAlgorithm), claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

func (j *JWTUtil) ParseToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.SecretKey), nil
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

func (j *JWTUtil) ParseRefreshToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.SecretKey), nil
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
