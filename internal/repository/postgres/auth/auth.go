package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user models.User) (int, error) {
	var id int
	row := r.db.QueryRow(createUserQuery, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUser(username, password string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, getUserQuery, username, password)

	return user, err
}
