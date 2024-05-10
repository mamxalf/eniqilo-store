package request

import "github.com/mamxalf/eniqilo-store/shared/validator"

type ProductSKUFilterParams struct {
	Limit    int    `validate:"omitempty,min=1" form:"limit" query:"limit"`
	Offset   int    `validate:"omitempty,min=0" form:"offset" query:"offset"`
	Name     string `validate:"omitempty,min=1" form:"name" query:"name"`
	Category string `validate:"omitempty,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'" form:"category" query:"category"`
	SKU      string `validate:"omitempty,min=1" form:"sku" query:"sku"`
	Price    string `validate:"omitempty,oneof='asc' 'desc'" form:"price" query:"price"`
	InStock  *bool  `validate:"omitempty" form:"inStock" query:"inStock"`
}

func (r *ProductSKUFilterParams) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}
