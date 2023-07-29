package product

import "time"

type (
	Product struct {
		ID        int32     `json:"id"`
		Name      string    `json:"name"`
		Price     int64     `json:"price"`
		Stock     int32     `json:"stock"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	AddProduct struct {
		Name  string `json:"name" validate:"required"`
		Price int64  `json:"price" validate:"required"`
		Stock int32  `json:"stock" validate:"required"`
	}

	UpdateProduct struct {
		ID    int32  `json:"id" validate:"required"`
		Name  string `json:"name" validate:"required"`
		Price int64  `json:"price" validate:"required"`
		Stock int32  `json:"stock" validate:"required"`
	}
)
