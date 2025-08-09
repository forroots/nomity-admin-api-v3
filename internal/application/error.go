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
		WithCode("internal_server_error"),
		WithMessage("大変申し訳ございません。システムに不具合が生じております。\n管理者にお問い合わせください。"),
		WithError(err),
	)
}

// 想定外のリクエストによるエラーが発生した場合に使用する（クライアントには固定のメッセージを返す）
func NewBadRequestError(err error) *AppError {
	return NewAppError(
		WithHttpStatus(http.StatusBadRequest),
		WithCode("bad_request"),
		WithMessage("リクエストの内容に誤りがあります。"),
		WithError(err),
	)
}
