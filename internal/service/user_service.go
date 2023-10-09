package service

import (
	"context"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/pkg/cache"
)

type userService struct {
	repo  repo.UserRepo
	cache cache.Cache
}

// UserLogin implements UserService.
func (us *userService) Login(ctx context.Context, user *model.UserLoginRequest) error {
	return nil
}

type UserService interface {
	Login(ctx context.Context, user *model.UserLoginRequest) error
}

func NewUserService(repo repo.UserRepo, cache cache.Cache) UserService {
	return &userService{repo: repo, cache: cache}
}
