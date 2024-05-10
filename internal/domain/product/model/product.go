package model

import "github.com/google/uuid"

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
