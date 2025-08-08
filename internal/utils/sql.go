package utils

import (
	"database/sql"
	"time"
)

func ToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
func FromNullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func ToNullInt32(i *int) sql.NullInt32 {
	if i != nil {
		return sql.NullInt32{Int32: int32(*i), Valid: true}
	}
	return sql.NullInt32{}
}
func FromNullInt32(ni sql.NullInt32) *int {
	if ni.Valid {
		v := int(ni.Int32)
		return &v
	}
	return nil
}

func ToNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{}
}
func FromNullTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
