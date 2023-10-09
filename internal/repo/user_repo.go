package repo

import (
	"echo-mangosteen/internal/common/data"
	"echo-mangosteen/internal/model"
)

type userRepo struct {
	*data.Data
}

// Save implements UserRepo.
func (*userRepo) Save(user *model.User) error {
	panic("unimplemented")
}

type UserRepo interface {
	Save(user *model.User) error
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{data}
}
