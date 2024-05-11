package request

import (
	// "github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/shared/validator"

	"github.com/google/uuid"
)

type InsertTransactionRequest struct {
	CustomerID     uuid.UUID       `json:"customerId" validate:"required"`
	ProductDetails []ProductDetail `json:"productDetails" validate:"required,min=1"`
	Paid           int             `json:"paid" validate:"required,min=1"`
	Change         int             `json:"change" validate:"required,min=0"`
}
type ProductDetail struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

func (r *InsertTransactionRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *InsertTransactionRequest) ToModel() (transaction model.InsertTransaction, err error) {
	transaction = model.InsertTransaction{
		CustomerID: r.CustomerID,
		Paid:       r.Paid,
		Change:     r.Change,
	}
	var productDetails []model.ProductDetail
	for _, detail := range r.ProductDetails {
		productDetails = append(productDetails, model.ProductDetail{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
		})
	}

	transaction.ProductDetails = productDetails
	return
}
