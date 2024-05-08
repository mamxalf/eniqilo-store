package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `db:"id"`
	StaffID     uuid.UUID `db:"staff_id"`
	Name        string    `db:"name"`
	SKU         string    `db:"sku"`
	Category    string    `db:"category"`
	ImageURL    string    `db:"imageurl"`
	Notes       string    `db:"notes"`
	Price       int       `db:"price"`
	Stock       int       `db:"stock"`
	Location    string    `db:"location"`
	IsAvailable bool      `db:"isavailable"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type InsertProduct struct {
	StaffID     uuid.UUID `db:"staff_id"`
	Name        string    `db:"name"`
	SKU         string    `db:"sku"`
	Category    string    `db:"category"`
	ImageURL    string    `db:"imageurl"`
	Notes       string    `db:"notes"`
	Price       int       `db:"price"`
	Stock       int       `db:"stock"`
	Location    string    `db:"location"`
	IsAvailable bool      `db:"isavailable"`
}
