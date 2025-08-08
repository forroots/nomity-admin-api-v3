package utils

import "time"

func BoolPtr(val bool) *bool {
	return &val
}

func StringPtr(val string) *string {
	return &val
}

func IntPtr(val int) *int {
	return &val
}

func Int64Ptr(val int64) *int64 {
	return &val
}

func TimePtr(val time.Time) *time.Time {
	return &val
}

func ToAnySlice[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
