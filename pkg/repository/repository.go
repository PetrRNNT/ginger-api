package repository

import (
	"github.com/jmoiron/sqlx"
	ginger_api "github.com/petrrnnt/ginger-api"
)

type Authorization interface {
	CreateUser(user ginger_api.User) (int, error)
	GetUser(username, password string) (ginger_api.User, error)
}

type PostList interface {
	Create(userId int, list ginger_api.PostList) (int, error)
	GetAll(userId int) ([]ginger_api.PostList, error)
	GetById(userId int, listId int) (ginger_api.PostList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input ginger_api.UpdateListInput) error
}

type PostItem interface {
	Create(listId int, item ginger_api.PostItem) (int, error)
	GetAll(userId int, listId int) ([]ginger_api.PostItem, error)
	GetById(userId int, itemId int) (ginger_api.PostItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input ginger_api.UpdateItemInput) error
}

type Repository struct {
	Authorization
	PostList
	PostItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PostList:      NewPostListPostgres(db),
		PostItem:      NewPostItemPostgres(db),
	}
}
