package todo_categories

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
)

type TodoCategories struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoCategories {
	return &TodoCategories{db: db}
}

func (r *TodoCategories) Create(userId int, categoryName string) (int, error) {
	var id int
	fmt.Println(userId)
	row := r.db.QueryRow(createCategoryQuery, userId, categoryName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TodoCategories) GetAll(userId int) ([]models.TodoCategory, error) {
	var categories []models.TodoCategory

	err := r.db.Select(&categories, getAllQuery, userId)

	return categories, err
}

func (r *TodoCategories) GetById(categoryId int) (models.TodoCategory, error) {
	var category models.TodoCategory

	err := r.db.Get(&category, getById, categoryId)

	return category, err
}

func (r *TodoCategories) Delete(userId int, categoryName string) error {
	// Вызов хранимой процедуры для удаления категории и связанных задач
	_, err := r.db.Exec(deleteQuery, userId, categoryName)
	return err
}

func (r *TodoCategories) Update(userId, categoryId int, input models.UpdateTodoCategory) error {
	// Создаем построитель запросов
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Начинаем с базового запроса UPDATE
	query := psql.Update("categories").Where("id = ?", categoryId)

	// Добавляем условия для обновления
	if input.Name != nil {
		query = query.Set("name", *input.Name)
	}

	// Добавляем условие для связи с пользователем
	query = query.Where("user_id = ?", userId)

	// Получаем финальный SQL запрос и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	logrus.Debugf("updateQuery: %s", sql)
	logrus.Debugf("args: %s", args)

	// Выполняем запрос
	_, err = r.db.Exec(sql, args...)
	return err
}
