package response

import (
	"time"

	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"imageurl"`
	Notes       string    `json:"notes"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isavailable"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
