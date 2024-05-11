package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID             uuid.UUID       `db:"id" json:"transactionId"`
	CustomerID     uuid.UUID       `db:"customer_id" json:"customerId"`
	ProductDetails []ProductDetail `db:"product_details" json:"productDetails"`
	Paid           int             `db:"paid" json:"paid"`
	Change         int             `db:"change" json:"change"`
	CreatedAt      time.Time       `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time       `db:"updated_at"`
}
type InsertTransaction struct {
	CustomerID     uuid.UUID       `db:"customer_id" json:"customerId"`
	ProductDetails []ProductDetail `db:"product_details" json:"productDetails"`
	Paid           int             `db:"paid" json:"paid"`
	Change         int             `db:"change" json:"change"`
}
type ProductDetail struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
