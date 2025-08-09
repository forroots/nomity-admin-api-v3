// internal/interface/common/cookie_settings.go
package common

import (
	"net/http"

	"github.com/forroots/nomity-admin-api-v3/internal/config"
)

type CookieSettings struct {
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

func NewCookieSettingsFromConfig(cfg config.CookieConfig) CookieSettings {
	sameSite := ConvertSameSite(cfg.SameSite)

	return CookieSettings{
		Name:     cfg.Name,
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
		SameSite: sameSite,
		MaxAge:   cfg.MaxAge,
	}
}
