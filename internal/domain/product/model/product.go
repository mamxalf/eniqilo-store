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

type ProductSKU struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Sku       string    `db:"sku" json:"sku"`
	Category  string    `db:"category" json:"category"`
	ImageUrl  string    `db:"image_url" json:"imageUrl"`
	Stock     int       `db:"stock" json:"stock"`
	Price     int       `db:"price" json:"price"`
	Location  string    `db:"location" json:"location"`
	CreatedAt string    `db:"created_at" json:"createdAt"`
	UpdatedAt string    `db:"updated_at" json:"updatedAt"`
}
