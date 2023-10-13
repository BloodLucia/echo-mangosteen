package service

import (
	"context"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/pkg/errors"
	"time"

	"github.com/spf13/cast"
)

type itemService struct {
	repo repo.ItemRepo
}

// CreateItem implements ItemService.
func (srv *itemService) CreateItem(ctx context.Context, userId string, req *model.CreateItemRequest) error {
	item := &model.Item{
		UserId:     cast.ToInt64(userId),
		TagId:      cast.ToInt64(req.TagId),
		Kind:       req.Kind,
		Amount:     req.Amount,
		HappenedAt: time.Now(),
	}

	err := srv.repo.Create(ctx, item)
	if err != nil {
		return errors.InternalServer().WithMsg("创建失败, 请稍后再试~")
	}

	return nil
}

type ItemService interface {
	CreateItem(ctx context.Context, userId string, req *model.CreateItemRequest) error
}

func NewItemService(repo repo.ItemRepo) ItemService {
	return &itemService{repo}
}
