package application

import "net/http"

type AppError struct {
	HttpStatus int
	Code       string
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

type AppErrorOption func(*AppError)

func NewAppError(opts ...AppErrorOption) *AppError {
	e := &AppError{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func WithHttpStatus(status int) AppErrorOption {
	return func(e *AppError) {
		e.HttpStatus = status
	}
}

func WithCode(code string) AppErrorOption {
	return func(e *AppError) {
		e.Code = code
	}
}

func WithMessage(msg string) AppErrorOption {
	return func(e *AppError) {
		e.Message = msg
	}
}

func WithError(err error) AppErrorOption {
	return func(e *AppError) {
		e.Err = err
	}
}

// 想定外のシステムエラーが発生した場合に使用する（クライアントには固定のメッセージを返す）
func NewInternalServerError(err error) *AppError {
	return NewAppError(
		WithHttpStatus(http.StatusInternalServerError),
		WithCode(APP_ERROR_INTERNAL_SERVER_ERROR),
		WithMessage("申し訳ございません。予期せぬエラーが発生しました。後ほど再度お試しください。"),
		WithError(err),
	)
}

// 想定外のリクエストによるエラーが発生した場合に使用する（クライアントには固定のメッセージを返す）
func NewBadRequestError(err error) *AppError {
	return NewAppError(
		WithHttpStatus(http.StatusBadRequest),
		WithCode(APP_ERROR_BAD_REQUEST),
		WithMessage("不正なリクエストです。入力内容を確認してください。"),
		WithError(err),
	)
}

const (
	APP_ERROR_INTERNAL_SERVER_ERROR = "internal_server_error"
	APP_ERROR_BAD_REQUEST           = "bad_request"
	APP_ERROR_SESSION_NOT_FOUND     = "session_not_found"
	APP_ERROR_UNAUTHORIZED          = "unauthorized"
)
