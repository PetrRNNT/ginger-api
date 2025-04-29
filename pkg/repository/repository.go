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
}

type PostItem interface {
}

type Repository struct {
	Authorization
	PostList
	PostItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
