package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
