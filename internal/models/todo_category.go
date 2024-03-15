package models

import "errors"

type TodoCategory struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name" binding:"required"`
	UserId int    `json:"user_id" db:"user_id"`
}

type UpdateTodoCategory struct {
	Name *string `json:"name" db:"name" binding:"required"`
}

func (i UpdateTodoCategory) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
