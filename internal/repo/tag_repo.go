package repo

import (
	"context"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/model"
	"log"
)

type tagRepo struct {
	*data.Data
}

// Create create a tag.
func (repo *tagRepo) Create(ctx context.Context, tag *model.Tag) error {
	if _, err := repo.Data.DB.Context(ctx).Insert(tag); err != nil {
		log.Panicln(err)
	}

	return nil
}

type TagRepo interface {
	Create(ctx context.Context, tag *model.Tag) error
}

func NewTagRepo(data *data.Data) TagRepo {
	return &tagRepo{Data: data}
}
