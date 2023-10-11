package repo

import (
	"context"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/model"
	"echo-mangosteen/pkg/errors"
	"log"
)

type tagRepo struct {
	*data.Data
}

// Deelte implements TagRepo.
func (repo *tagRepo) Delete(ctx context.Context, tagId string) error {
	tag := &model.Tag{}
	exist, err := repo.Data.DB.Context(ctx).Where("id = ?", tagId).Get(tag)
	if err != nil {
		return err
	}

	if !exist {
		return err
	}

	if _, err := repo.Data.DB.Context(ctx).Delete(tag); err != nil {
		return errors.InternalServer().WithMsg("删除标签失败")
	}

	return nil
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
	Delete(ctx context.Context, tagId string) error
}

func NewTagRepo(data *data.Data) TagRepo {
	return &tagRepo{Data: data}
}
