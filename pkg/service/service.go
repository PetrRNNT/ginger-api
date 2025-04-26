package service

import "github.com/petrrnnt/ginger-api/pkg/repository"

type Authorization interface {
}

type PostList interface {
}

type PostItem interface {
}

type Service struct {
	Authorization
	PostList
	PostItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
