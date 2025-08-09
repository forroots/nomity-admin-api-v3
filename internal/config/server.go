package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port           int      `mapstructure:"port"`
	BasePath       string   `mapstructure:"base_path"`
	TrustedProxies []string `mapstructure:"trusted_proxies"`
}

type CORSConfig struct {
	Enabled              bool     `mapstructure:"enabled"`
	Logging              bool     `mapstructure:"logging"`
	AllowOrigins         []string `mapstructure:"allow_origins"`
	AdvancedAllowHeaders []string `mapstructure:"advanced_allow_headers"`
}

type CSRFConfig struct {
	Enabled         bool   `mapstructure:"enabled"`
	TokenCookieName string `mapstructure:"token_cookie_name"`
	TokenHeader     string `mapstructure:"token_header"`
}

type CookieConfig struct {
	Name     string `mapstructure:"name"`
	Path     string `mapstructure:"path"`
	Domain   string `mapstructure:"domain"`
	Secure   bool   `mapstructure:"secure"`
	HttpOnly bool   `mapstructure:"http_only"`
	SameSite string `mapstructure:"same_site"`
	MaxAge   int    `mapstructure:"max_age"`
}

// SetCookie
func (conf *CookieConfig) SetCookie(c *gin.Context, value string) {
	c.SetSameSite(convertSameSite(conf.SameSite))
	c.SetCookie(
		conf.Name,
		value,
		conf.MaxAge,
		conf.Path,
		conf.Domain,
		conf.Secure,
		conf.HttpOnly,
	)
}

// GetCookie
func (conf *CookieConfig) GetCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(conf.Name)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

// DeleteCookie
func (conf *CookieConfig) DeleteCookie(c *gin.Context) {
	c.SetCookie(
		conf.Name,
		"",
		-1,
		conf.Path,
		conf.Domain,
		conf.Secure,
		conf.HttpOnly,
	)
}

func convertSameSite(s string) http.SameSite {
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
