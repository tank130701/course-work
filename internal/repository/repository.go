package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/auth"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/todo_item"
	"github.com/tank130701/course-work/todo-app/back-end/internal/repository/postgres/todo_list"
)

type IAuthorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type ITodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type ITodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetList(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization IAuthorization
	TodoList      ITodoList
	TodoItem      ITodoItem
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuth(db),
		TodoList:      todo_list.NewTodoListPostgres(db),
		TodoItem:      todo_item.NewTodoItemPostgres(db),
	}
}
