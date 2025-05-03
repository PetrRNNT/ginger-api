package service

import (
	ginger_api "github.com/petrrnnt/ginger-api"
	"github.com/petrrnnt/ginger-api/pkg/repository"
)

type PostListService struct {
	repo repository.PostList
}

func NewPostListService(repo repository.PostList) *PostListService {
	return &PostListService{repo: repo}
}

func (s *PostListService) Create(userId int, list ginger_api.PostList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *PostListService) GetAll(userId int) ([]ginger_api.PostList, error) {
	return s.repo.GetAll(userId)
}

func (s *PostListService) GetById(userId int, listId int) (ginger_api.PostList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *PostListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *PostListService) Update(userId int, listId int, input ginger_api.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
