package service

import "echo-mangosteen/internal/repo"

type tagService struct {
	repo repo.TagRepo
}

type TagService interface {
}

func NewTagService(repo repo.TagRepo) TagService {
	return &tagService{repo}
}
