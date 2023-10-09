package repo

import (
	"context"

	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/model"
)

type userRepo struct {
	*data.Data
}

// FindOrCreate find a user by email and create if doesn't exist.
func (ur *userRepo) FindOrCreateByEmail(ctx context.Context, user *model.User) error {
	exist, err := ur.Data.DB.Context(ctx).Where("email = ?", user.Email).Get(user)
	if err != nil {
		return nil
	}
	if exist {
		if _, err := ur.Data.DB.Insert(user); err != nil {
			return err
		}
		return nil
	}
	panic("unimplemented")
}

type UserRepo interface {
	FindOrCreateByEmail(ctx context.Context, user *model.User) error
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{data}
}
