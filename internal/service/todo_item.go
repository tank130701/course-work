package service

import (
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository"
)

type TodoItemService struct {
	repo     repository.ITodoItem
	listRepo repository.ITodoList
}

func NewTodoItemService(repo repository.ITodoItem, listRepo repository.ITodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item models.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	return s.repo.GetList(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
