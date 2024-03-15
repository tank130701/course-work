package service

import (
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoCategory interface {
	Create(userId int, categoryName string) (int, error)
	GetAll(userId int) ([]models.TodoCategory, error)
	GetById(categoryId int) (models.TodoCategory, error)
	Delete(userId int, categoryName string) error
	Update(userId, listId int, input models.UpdateTodoCategory) error
}

type TodoItem interface {
	Create(userId, categoryId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoCategory
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoCategory:  NewTodoCategoriesService(repos.TodoCategories),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoCategories),
	}
}
