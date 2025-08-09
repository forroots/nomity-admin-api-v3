// internal/interface/common/cookie_settings.go
package common

import (
	"net/http"

	"github.com/forroots/nomity-admin-api-v3/internal/config"
	"github.com/gin-gonic/gin"
)

type CookieConfig struct {
	Name     string
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
	MaxAge   int
}

func ConvertSameSite(s string) http.SameSite {
	switch s {
	case "Lax":
		return http.SameSiteLaxMode
	case "Strict":
		return http.SameSiteStrictMode
	case "None":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}

func NewCookieSettingsFromConfig(cfg config.CookieConfig) CookieConfig {
	sameSite := ConvertSameSite(cfg.SameSite)

	return CookieConfig{
		Name:     cfg.Name,
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
		SameSite: sameSite,
		MaxAge:   cfg.MaxAge,
	}
}

func (cs *CookieConfig) SetCookie(c *gin.Context, value string) {
	c.SetCookie(
		cs.Name,
		value,
		cs.MaxAge,
		cs.Path,
		cs.Domain,
		cs.Secure,
		cs.HttpOnly,
	)
}

func (cs *CookieConfig) GetCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(cs.Name)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func (cs *CookieConfig) DeleteCookie(c *gin.Context) {
	c.SetCookie(
		cs.Name,
		"",
		-1,
		cs.Path,
		cs.Domain,
		cs.Secure,
		cs.HttpOnly,
	)
}
