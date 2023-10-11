package repo

import "echo-mangosteen/internal/data"

type tagRepo struct {
	*data.Data
}

type TagRepo interface {
}

func NewTagRepo(data *data.Data) TagRepo {
	return &tagRepo{Data: data}
}
