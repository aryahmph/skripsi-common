package http

import "time"

type (
	Response struct {
		Data any `json:"data"`
	}

	User struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		PasswordHash string    `json:"password_hash"`
		Role         string    `json:"role"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	AddUser struct {
		Name     string `json:"name" validate:"required,lte=255"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6,lte=16"`
	}

	Login struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6,lte=16"`
	}
)
