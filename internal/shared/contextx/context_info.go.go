package contextx

import (
	"context"
	"log/slog"
)

// ContextInfo はコンテキストにセットする追加情報などを保持する構造体です。
type ContextInfo struct {
	AdminUserID *int64 // 認証後に詰める用
	Logger      *slog.Logger
	RequestID   string
	SessionID   string
	IPAddress   string
	Method      string
	Path        string
}

type contextInfoKey struct{}

func WithContextInfo(ctx context.Context, info *ContextInfo) context.Context {
	return context.WithValue(ctx, contextInfoKey{}, info)
}

func GetContextInfo(ctx context.Context) (*ContextInfo, bool) {
	info, ok := ctx.Value(contextInfoKey{}).(*ContextInfo)
	return info, ok
}

func GetLogger(ctx context.Context) *slog.Logger {
	// enriched logger があれば使う、なければdefaultのslogを使う
	if info, ok := GetContextInfo(ctx); ok && info.Logger != nil {
		return info.Logger
	}
	return slog.Default()
}

func GetAdminUserID(ctx context.Context) *int64 {
	if info, ok := GetContextInfo(ctx); ok {
		return info.AdminUserID
	}
	return nil
}

func SetAdminUserID(ctx context.Context, userID int64) {
	if info, ok := GetContextInfo(ctx); ok {
		info.AdminUserID = &userID
	}
}
