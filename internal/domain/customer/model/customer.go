package model

import (
	"github.com/google/uuid"
	"time"
)

type CustomerRegister struct {
	Name  string `db:"name"`
	Phone string `db:"phone"`
}

type Customer struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
