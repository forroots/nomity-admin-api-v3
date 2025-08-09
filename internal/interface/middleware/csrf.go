package middleware

import (
	"log/slog"
	"net/http"

	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/forroots/nomity-admin-api-v3/internal/utils"
	"github.com/gin-gonic/gin"
)

func NewCSRFHandler(cookieName, headerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 安全なメソッドは対象外
		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodHead || c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// enriched logger があれば使う、なければ defaultのslogを使う
		logger := getLogger(c)

		errCode := "csrf_error"

		// CSRFトークンをクッキーから取得
		cookieToken, err := c.Cookie(cookieName)
		if err != nil {
			logger.Error("CSRF cookie not found", slog.Any("error", err))
			response.Error(
				c,
				http.StatusBadRequest,
				errCode,
				"MissingCSRFCookie",
				"CSRFクッキーが見つかりませんでした",
			)
			return
		}
		if cookieToken == "" {
			logger.Error("CSRF cookie is empty")
			response.Error(
				c,
				http.StatusBadRequest,
				errCode,
				"EmptyCSRFCookie",
				"CSRFクッキーが空です",
			)
			return
		}

		// CSRFトークンをヘッダーから取得
		headerToken := c.GetHeader(headerName)
		if headerToken == "" {
			logger.Error("CSRF header not found", slog.String("header", headerName))
			response.Error(
				c,
				http.StatusBadRequest,
				errCode,
				"MissingCSRFHeader",
				"CSRFトークンヘッダーが見つかりませんでした",
			)
			return
		}

		if headerToken != cookieToken {
			logger.Error("CSRF token mismatch", slog.String("cookie", cookieToken), slog.String("header", headerToken))
			response.Error(
				c,
				http.StatusForbidden,
				errCode,
				"CSRFTokenMismatch",
				"CSRFトークンが一致しません",
			)
			return
		}

		c.Next()
	}
}

func SetCSRFCookieIfNotExists(cookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if c.Request.Method != http.MethodGet {
		// 	// GETリクエストでなければCSRFトークンの生成、セットは行わない
		// 	c.Next()
		// 	return
		// }

		// _, err := c.Cookie(cookieName)
		// if err == nil {
		// 	// 既にCSRFトークンがクッキーに存在する場合は何もしない
		// 	c.Next()
		// 	return
		// }

		// すでにCSRFトークンがクッキーに存在する場合は期限をスライドする
		cookeToken, err := c.Cookie(cookieName)
		if err == nil && cookeToken != "" {
			// クッキーの期限をスライドする
			c.SetCookie(cookieName, cookeToken, 3600, "/", "", true, false)
			c.Next()
			return
		}

		// トークン生成(1時間有効)
		token := utils.NewCryptoULIDString()
		c.SetCookie(cookieName, token, 3600, "/", "", true, false)
		c.Next()
	}
}
