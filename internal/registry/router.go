package registry

import (
	"github.com/forroots/nomity-admin-api-v3/internal/interface/common"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/handler"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Registry) NewRouter() *gin.Engine {
	if r.Config.EnvProduction() {
		gin.SetMode(gin.ReleaseMode) // 本番環境ではReleaseモード
	}
	router := gin.New()

	if err := router.SetTrustedProxies(r.Config.Server.TrustedProxies); err != nil {
		r.Logger.Error("Failed to set trusted proxies", "error", err)
	}

	router.Use(middleware.EnrichContext(r.Logger, r.Config.Cookie.SessionID.Name))                                               // コンテキストにリクエスト情報とカスタムロガーを追加
	router.Use(middleware.RequestLogger())                                                                                       // 共通ログ出力ハンドラ
	router.Use(middleware.CustomRecovery())                                                                                      // panicハンドラ
	router.Use(middleware.NewErrorHandler())                                                                                     // エラーハンドラ
	router.Use(middleware.NewCORSWithLoggingHandler(r.Config.CORS.AllowOrigins, r.Config.CORS.AdvancedAllowHeaders))             // CORSハンドラ
	router.Use(middleware.NewCORSHandler(r.Config.CORS.AllowOrigins, r.Config.CORS.AdvancedAllowHeaders, r.Config.CORS.Logging)) // CORSハンドラ
	if r.Config.CSRF.Enabled {
		router.Use(middleware.SetCSRFCookieIfNotExists(common.NewCookieSettingsFromConfig(r.Config.Cookie.CSRFToken))) // CSRFトークンをクッキーにセットするハンドラ
		router.Use(middleware.NewCSRFHandler(r.Config.CSRF.TokenCookieName, r.Config.CSRF.TokenHeader))                // CSRFトークンハンドラ
	} else {
		r.Logger.Warn("CSRF protection is disabled")
	}

	api := router.Group(r.Config.Server.BasePath)
	api.GET("/health", func(ctx *gin.Context) {}) // ヘルスチェックエンドポイント

	// routes of "test"
	r.SetupTestRoutes(api.Group("/test"))

	return router
}

func (r *Registry) SetupTestRoutes(router *gin.RouterGroup) {
	h := handler.NewMailHandler(r.Mailer)
	router.GET("/panic", h.PanicTest)              // パニックテスト用エンドポイント
	router.GET("/error", h.ErrorTest)              // エラーテスト用エンドポイント
	router.GET("/app-error", h.AppErrorTest)       // エラーテスト用エンドポイント
	router.GET("/send-test-mail", h.SendTestEmail) // メール送信テスト用エンドポイント
	router.POST("just-post", h.JustPost)           // POSTメソッドのテスト用エンドポイント
}
