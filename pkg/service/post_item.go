package service

import (
	ginger_api "github.com/petrrnnt/ginger-api"
	"github.com/petrrnnt/ginger-api/pkg/repository"
)

type PostItemService struct {
	repo     repository.PostItem
	listRepo repository.PostList
}

func NewPostItemService(repo repository.PostItem, listRepo repository.PostList) *PostItemService {
	return &PostItemService{repo: repo, listRepo: listRepo}
}

func (s *PostItemService) Create(userId int, listId int, item ginger_api.PostItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *PostItemService) GetAll(userId int, listId int) ([]ginger_api.PostItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *PostItemService) GetById(userId int, itemId int) (ginger_api.PostItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *PostItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *PostItemService) Update(userId int, itemId int, input ginger_api.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
