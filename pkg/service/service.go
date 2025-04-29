package service

import (
	ginger_api "github.com/petrrnnt/ginger-api"
	"github.com/petrrnnt/ginger-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user ginger_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
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
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
