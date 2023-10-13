package service

import (
	"context"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/pkg/errors"

	"github.com/spf13/cast"
)

type tagService struct {
	repo repo.TagRepo
}

// AddTag implements TagService.
func (srv *tagService) AddTag(ctx context.Context, userId string, req *model.TagAddRequest) error {
	tag := &model.Tag{
		UserId: cast.ToInt64(userId),
		Name:   req.Name,
		Type:   req.Type,
		Sign:   req.Sign,
	}
	if err := srv.repo.Create(ctx, tag); err != nil {
		return errors.InternalServer().WithMsg("添加标签失败, 请稍后再试")
	}

	return nil
}

type TagService interface {
	AddTag(ctx context.Context, userId string, req *model.TagAddRequest) error
}

func NewTagService(repo repo.TagRepo) TagService {
	return &tagService{repo}
}
