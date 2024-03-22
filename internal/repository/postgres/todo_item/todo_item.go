package todo_item

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
)

type TodoItemPostgres struct {
	db *sqlx.DB
	qb *squirrel.StatementBuilderType
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &TodoItemPostgres{db: db, qb: &qb}
}

func (r *TodoItemPostgres) Create(userId, categoryId int, item models.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskId int

	row := tx.QueryRow(createTaskQuery, item.Title, item.Description, item.Status, userId, categoryId)
	err = row.Scan(&taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskId, tx.Commit()
}

func (r *TodoItemPostgres) GetList(userId, categoryId int) ([]models.TodoItem, error) {
	var items []models.TodoItem

	if err := r.db.Select(&items, getAllQuery, userId, categoryId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	if err := r.db.Get(&item, getByIdQuery, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	//_, err := r.db.Exec(deleteQuery, userId, itemId)
	_, err := r.db.Exec(deleteQuery, itemId)
	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input models.UpdateItemInput) error {
	// Start building the query
	query := r.qb.Update("tasks")

	// Add conditions to the query
	if input.Title != nil {
		query = query.Set("title", *input.Title)
	}

	if input.Description != nil {
		query = query.Set("description", *input.Description)
	}

	if input.Status != nil {
		query = query.Set("status", *input.Status)
	}

	// Finalize the query
	queryRes, args, err := query.
		//Where(squirrel.Expr("id IN (SELECT item_id FROM "+"items"+" WHERE list_id IN (SELECT list_id FROM users WHERE user_id = ?))", userId)).
		Where("id = ?", itemId).
		ToSql()

	if err != nil {
		return err
	}

	// Execute the query
	_, err = r.db.Exec(queryRes, args...)
	return err
}
