package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,min=3"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=3"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
