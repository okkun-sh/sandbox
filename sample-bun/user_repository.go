package sample_bun

import (
	"context"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	ctx context.Context
	db  bun.IDB
}

type User struct {
	Id   int64
	Name string
}

func NewUserRepository(ctx context.Context, db bun.IDB) *UserRepository {
	return &UserRepository{
		ctx: ctx,
		db:  db,
	}
}

func (u *UserRepository) FindUserById(id int64) (*User, error) {
	user := &User{}
	err := u.db.NewSelect().Model(user).Where("id = ?", id).Scan(u.ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FindByName(name string) (*User, error) {
	user := &User{}
	err := u.db.NewSelect().Model(user).Where("name = ?", name).Scan(u.ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
