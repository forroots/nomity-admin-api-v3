package contextx

import (
	"context"
	"log/slog"
)

// ContextInfo はコンテキストにセットする追加情報などを保持する構造体です。
type ContextInfo struct {
	Logger    *slog.Logger
	RequestID string
	SessionID string
	IPAddress string
	Method    string
	Path      string
}

type contextInfoKey struct{}

func WithContextInfo(ctx context.Context, info *ContextInfo) context.Context {
	return context.WithValue(ctx, contextInfoKey{}, info)
}

func GetContextInfo(ctx context.Context) (*ContextInfo, bool) {
	info, ok := ctx.Value(contextInfoKey{}).(*ContextInfo)
	return info, ok
}
