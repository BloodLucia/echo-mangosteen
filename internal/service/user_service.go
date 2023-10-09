package service

import "echo-mangosteen/internal/repo"

type userService struct {
	repo repo.UserRepo
}

// UserLogin implements UserService.
func (us *userService) UserLogin() error {
	panic("unimplemented")
}

type UserService interface {
	UserLogin() error
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{repo: repo}
}
