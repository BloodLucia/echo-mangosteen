package repo

import "echo-mangosteen/internal/data"

type itemRepo struct {
	*data.Data
}

type ItemRepo interface {
}

func NewItemRepo(data *data.Data) ItemRepo {
	return &itemRepo{data}
}
