package request

import (
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/shared/validator"

	"github.com/google/uuid"
)

type InsertProductRequest struct {
	StaffID     uuid.UUID `json:"-"`
	Name        string    `validate:"required,min=1,max=30"`
	SKU         string    `validate:"required,min=1,max=30"`
	Category    string    `validate:"required,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'"`
	ImageURL    string    `validate:"required,url"`
	Notes       string    `validate:"required,min=1,max=200"`
	Price       int       `validate:"required,min=1"`
	Stock       int       `validate:"required,min=0,max=100000"`
	Location    string    `validate:"required,min=1,max=200"`
	IsAvailable bool      `validate:"required"`
}

func (r *InsertProductRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *InsertProductRequest) ToModel() (product model.InsertProduct, err error) {
	product = model.InsertProduct{
		StaffID:     r.StaffID,
		Name:        r.Name,
		SKU:         r.SKU,
		Category:    r.Category,
		ImageURL:    r.ImageURL,
		Notes:       r.Notes,
		Price:       r.Price,
		Stock:       r.Stock,
		Location:    r.Location,
		IsAvailable: r.IsAvailable,
	}
	return
}

type UpdateProductRequest struct {
	Name        string `validate:"required"`
	SKU         string `validate:"required"`
	Category    string `validate:"required,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'"`
	ImageURL    string `validate:"required,url"`
	Notes       string `validate:"required,min=1,max=200"`
	Price       int    `validate:"required,min=1"`
	Stock       int    `validate:"required,min=0,max=100000"`
	Location    string `validate:"required,min=1,max=200"`
	IsAvailable bool   `validate:"required"`
}

func (r *UpdateProductRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *UpdateProductRequest) ToModel() (product model.Product, err error) {
	product = model.Product{
		Name:        r.Name,
		SKU:         r.SKU,
		Category:    r.Category,
		ImageURL:    r.ImageURL,
		Notes:       r.Notes,
		Price:       r.Price,
		Stock:       r.Stock,
		Location:    r.Location,
		IsAvailable: r.IsAvailable,
	}
	return
}

type ProductQueryParams struct {
	ID          string `json:"id" validate:"omitempty,uuid"`
	Limit       int    `json:"limit" validate:"omitempty,min=1"`
	Offset      int    `json:"offset" validate:"omitempty,min=0"`
	Name        string `json:"name" validate:"omitempty"`
	IsAvailable bool   `json:"isAvailable" validate:"-"`
	Category    string `json:"category" validate:"omitempty,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'"`
	SKU         string `json:"sku" validate:"omitempty"`
	Price       string `json:"price" validate:"omitempty,oneof='asc' 'desc'"`
	InStock     bool   `json:"inStock" validate:"-"`
	Owned       bool   `json:"owned" validate:"omitempty"`
	Search      string `json:"search" validate:"omitempty,min=1"`
	CreatedAt   string `json:"createdAt" validate:"omitempty,oneof='asc' 'desc'"`
}

func (r *ProductQueryParams) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}
