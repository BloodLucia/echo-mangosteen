package service

import "echo-mangosteen/internal/repo"

type itemService struct {
	repo repo.ItemRepo
}

type ItemService interface {
}

func NewItemService(repo repo.ItemRepo) ItemService {
	return &itemService{repo}
}
