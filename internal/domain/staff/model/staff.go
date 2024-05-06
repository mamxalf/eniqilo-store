package model

import (
	"github.com/google/uuid"
	"time"
)

type StaffRegister struct {
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
}

type Staff struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
