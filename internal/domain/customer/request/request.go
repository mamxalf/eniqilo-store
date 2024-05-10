package request

import (
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/model"
	"github.com/mamxalf/eniqilo-store/shared/validator"
)

type RegisterRequest struct {
	PhoneNumber string `validate:"required,min=10,max=16,e164" json:"phoneNumber" example:"+6285123546789"`
	Name        string `validate:"required,min=5,max=50" json:"name" example:"Yuri co"`
}

func (r *RegisterRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *RegisterRequest) ToModel() (register model.CustomerRegister, err error) {
	register = model.CustomerRegister{
		Name:  r.Name,
		Phone: r.PhoneNumber,
	}
	return
}
