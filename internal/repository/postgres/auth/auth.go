package auth

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tank130701/course-work/todo-app/back-end/internal/errs"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
)

const usernameConstraint = "users_username_key"

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user models.User) (int, error) {
	var id int
	row := r.db.QueryRow(createUserQuery, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		var dbErr *pq.Error
		if errors.As(err, &dbErr) && dbErr.Constraint == usernameConstraint {
			return 0, errs.NewErrorUsernameAlreadyExists(user.Username)
		}
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUser(username, password string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, getUserQuery, username, password)
	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, errs.NewErrorNotFound(username)
	}

	return user, nil
}
