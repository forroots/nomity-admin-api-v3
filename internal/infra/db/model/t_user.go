package model

import (
	"time"
)

// TUser maps to public.t_users table.
type TUser struct {
	ID           int64     `bun:"id,pk"`
	CreatedAt    time.Time `bun:"created_at,nullzero"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero"`
	DeletedAt    time.Time `bun:"deleted_at,soft_delete,nullzero"`
	ParentUserID *int64    `bun:"parent_user_id"`
	Name         string    `bun:"name,notnull"`
	Email        *string   `bun:"email,nullzero"`
	PasswordHash string    `bun:"password_hash,notnull"`
	SerchingID   *string   `bun:"serching_id,nullzero"`
	AllowExpose  bool      `bun:"allow_expose,notnull,default:false"`
	RoleCD       string    `bun:"role_cd,notnull"`
}
