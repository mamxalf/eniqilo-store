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
	ID        uuid.UUID `db:"id" json:"userId"`
	Name      string    `db:"name" json:"name"`
	Phone     string    `db:"phone" json:"phoneNumber"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
