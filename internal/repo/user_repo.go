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
func (ur *userRepo) FindOrCreateByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	exist, err := ur.Data.DB.Context(ctx).Where("email = ?", user.Email).Get(user)
	if err != nil {
		return nil, err
	}
	if !exist {
		// 创建用户
		user.Email = email
		if _, err := ur.Data.DB.Insert(user); err != nil {
			return nil, err
		}
		return user, nil
	}
	return user, nil
}

type UserRepo interface {
	FindOrCreateByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{data}
}
