package std

import (
	"encoding/json"
	"regexp"

	"github.com/forroots/nomity-admin-api-v3/pkg/vo"
)

type Email string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func NewEmail(s string) (Email, error) {
	if s == "" {
		return "", vo.NewVOErrorf("email cannot be empty")
	}
	if !emailRegex.MatchString(s) {
		return "", vo.NewVOErrorf("invalid email format '%s'", s)
	}
	return Email(s), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}
