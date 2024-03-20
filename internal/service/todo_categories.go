package service

import (
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository"
)

type TodoCategoriesService struct {
	repo repository.ITodoCategories
}

func NewTodoCategoriesService(repo repository.ITodoCategories) *TodoCategoriesService {
	return &TodoCategoriesService{repo: repo}
}

func (s *TodoCategoriesService) Create(userId int, categoryName string) (int, error) {
	return s.repo.Create(userId, categoryName)
}

func (s *TodoCategoriesService) GetAll(userId int) ([]models.TodoCategory, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoCategoriesService) GetById(categoryId int) (models.TodoCategory, error) {
	return s.repo.GetById(categoryId)
}

func (s *TodoCategoriesService) Delete(userId int, categoryId int) error {
	return s.repo.Delete(userId, categoryId)
}

func (s *TodoCategoriesService) Update(userId, listId int, input models.UpdateTodoCategory) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
