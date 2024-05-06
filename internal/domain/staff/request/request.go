package request

import (
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/model"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/utils"
	"github.com/mamxalf/eniqilo-store/shared/validator"
)

type RegisterRequest struct {
	PhoneNumber string `validate:"required,min=10,max=16,e164" json:"phoneNumber" example:"+6285123546789"`
	Name        string `validate:"required,min=5,max=50" json:"name" example:"Yuri co"`
	Password    string `validate:"required,alphanum,min=5,max=15" json:"password,omitempty" example:"s3Cr3Tk3y"`
}

func (r *RegisterRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *RegisterRequest) ToModel() (register model.StaffRegister, err error) {
	hashPassword, err := utils.HashPassword(r.Password)
	if err != nil {
		logger.ErrorWithMessage(err, "failed to hash password")
		return
	}
	register = model.StaffRegister{
		Name:     r.Name,
		Phone:    r.PhoneNumber,
		Password: hashPassword,
	}
	return
}
