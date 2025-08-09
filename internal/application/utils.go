package application

import "time"

func Deref(s *string, fallback string) string {
	if s != nil {
		return *s
	}
	return fallback
}

func DerefInt(i *int, fallback int) int {
	if i != nil {
		return *i
	}
	return fallback
}

func FormatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
