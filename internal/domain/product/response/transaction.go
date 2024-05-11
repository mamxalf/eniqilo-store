package response

import (
	"time"

	"github.com/google/uuid"
)

type TransactionResponse struct {
	TransactionID  uuid.UUID       `json:"transactionId"`
	CustomerID     uuid.UUID       `json:"customerId"`
	ProductDetails []ProductDetail `json:"productDetails"`
	Paid           int             `json:"paid"`
	Change         int             `json:"change"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type ProductDetail struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
