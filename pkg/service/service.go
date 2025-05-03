package service

import (
	ginger_api "github.com/petrrnnt/ginger-api"
	"github.com/petrrnnt/ginger-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user ginger_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type PostList interface {
	Create(userId int, list ginger_api.PostList) (int, error)
	GetAll(userId int) ([]ginger_api.PostList, error)
	GetById(userId int, listId int) (ginger_api.PostList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input ginger_api.UpdateListInput) error
}

type PostItem interface {
	Create(userId int, listId int, item ginger_api.PostItem) (int, error)
	GetAll(userId int, listId int) ([]ginger_api.PostItem, error)
	GetById(userId int, itemId int) (ginger_api.PostItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input ginger_api.UpdateItemInput) error
}

type Service struct {
	Authorization
	PostList
	PostItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostList:      NewPostListService(repos.PostList),
		PostItem:      NewPostItemService(repos.PostItem, repos.PostList),
	}
}
