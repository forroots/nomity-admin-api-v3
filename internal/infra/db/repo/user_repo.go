package repo

import (
	"context"

	"github.com/forroots/nomity-admin-api-v3/internal/infra/db/model"
	"github.com/uptrace/bun"
)

type UserRepo struct {
	db bun.IDB
}

func NewUserRepo(db bun.IDB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) WithTx(tx bun.Tx) *UserRepo {
	return &UserRepo{db: tx}
}

// ListAll selects all users (excluding soft-deleted).
func (r *UserRepo) ListAll(ctx context.Context) ([]model.TUser, error) {
	var users []model.TUser
	err := r.db.NewSelect().
		Model(&users).
		Order("id ASC").
		Scan(ctx)
	return users, err
}

// FindByID selects a single user by ID.
func (r *UserRepo) FindByID(ctx context.Context, id int64) (*model.TUser, error) {
	var user model.TUser
	err := r.db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
