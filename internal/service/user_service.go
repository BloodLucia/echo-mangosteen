package service

import (
	"context"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/pkg/cache"
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/jwt"
	"time"
)

type userService struct {
	repo  repo.UserRepo
	cache cache.Cache
	jwt   *jwt.JWT
}

// UserLogin implements UserService.
func (us *userService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	if req.Code != "123456" {
		return nil, errors.BadRequest().WithMsg("invalid validation code")
	}
	user, err := us.repo.FindOrCreateByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	// jwt := jwt.New()
	token, err := us.jwt.BuildToken(user.GetStringID(), time.Now().Add(10*time.Minute))
	if err != nil {
		return nil, errors.InternalServer().WithErr(err)
	}
	resp := &model.UserLoginResponse{Token: token}

	return resp, nil
}

type UserService interface {
	Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error)
}

func NewUserService(repo repo.UserRepo, cache cache.Cache, jwt *jwt.JWT) UserService {
	return &userService{repo: repo, cache: cache, jwt: jwt}
}
