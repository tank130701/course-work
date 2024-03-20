package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/auth"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/todo_categories"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/todo_item"
)

type IAuthorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type ITodoCategories interface {
	Create(userId int, categoryName string) (int, error)
	GetAll(userId int) ([]models.TodoCategory, error)
	GetById(categoryId int) (models.TodoCategory, error)
	Delete(userId int, categoryId int) error
	Update(userId, listId int, input models.UpdateTodoCategory) error
}

type ITodoItem interface {
	Create(userId, categoryId int, item models.TodoItem) (int, error)
	GetList(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization  IAuthorization
	TodoCategories ITodoCategories
	TodoItem       ITodoItem
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  auth.NewAuth(db),
		TodoCategories: todo_categories.NewTodoListPostgres(db),
		TodoItem:       todo_item.NewTodoItemPostgres(db),
	}
}
