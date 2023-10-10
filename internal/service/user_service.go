package service

import (
	"context"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/pkg/cache"
	"echo-mangosteen/pkg/errors"
)

type userService struct {
	repo  repo.UserRepo
	cache cache.Cache
}

// UserLogin implements UserService.
func (us *userService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	if req.Code != "123456" {
		return nil, errors.BadRequest().WithMsg("invalid validation code")
	}
	if err := us.repo.FindOrCreateByEmail(ctx, &model.User{Email: req.Email}); err != nil {
		return nil, errors.InternalServer()
	}
	// TODO 生成token
	return nil, nil
}

type UserService interface {
	Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error)
}

func NewUserService(repo repo.UserRepo, cache cache.Cache) UserService {
	return &userService{repo: repo, cache: cache}
}
