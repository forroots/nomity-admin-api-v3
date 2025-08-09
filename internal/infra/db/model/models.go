package model

import (
	"time"
)

type TUser struct {
	ID           int64      `bun:"id,pk"`
	CreatedAt    *time.Time `bun:"created_at"`
	UpdatedAt    *time.Time `bun:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete"`
	ParentUserID *int64     `bun:"parent_user_id"`
	Name         string     `bun:"name,notnull"`
	Email        *string    `bun:"email"`
	PasswordHash string     `bun:"password_hash,notnull"`
	SerchingID   *string    `bun:"serching_id"`
	AllowExpose  bool       `bun:"allow_expose,notnull,default:false"`
	RoleCD       string     `bun:"role_cd,notnull"`
}

type TAdminUser struct {
	ID               int64      `bun:"id,pk"`
	CreatedAt        *time.Time `bun:"created_at"`
	UpdatedAt        *time.Time `bun:"updated_at"`
	DeletedAt        *time.Time `bun:"deleted_at,soft_delete"`
	Name             string     `bun:"name,notnull"`
	Email            *string    `bun:"email"`
	PasswordHash     string     `bun:"password_hash,notnull"`
	EmailConfirmedAt *time.Time `bun:"email_confirmed_at"`
	Active           bool       `bun:"active,notnull,default:false"`
	Role             string     `bun:"role,notnull"`
}
