package handler

import (
	"fmt"

	"github.com/forroots/nomity-admin-api-v3/internal/application"
)

func NewSessionNotFoundError() *application.AppError {
	return application.NewBadRequestError(fmt.Errorf("セッションIDが取得できないもしくは空文字"))
}
