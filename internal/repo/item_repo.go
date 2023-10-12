package repo

import (
	"context"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/pkg/errors"
)

type itemRepo struct {
	*data.Data
}

// Create implements ItemRepo.
func (repo *itemRepo) Create(ctx context.Context, item *model.Item) error {
	if _, err := repo.Data.DB.Context(ctx).Insert(item); err != nil {
		return errors.InternalServer()
	}

	return nil
}

type ItemRepo interface {
	Create(ctx context.Context, item *model.Item) error
}

func NewItemRepo(data *data.Data) ItemRepo {
	return &itemRepo{data}
}
