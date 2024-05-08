package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/model"
)

type StaffRepository interface {
	Register(ctx context.Context, staffRegister *model.StaffRegister) (id uuid.UUID, err error)
	FindByPhone(ctx context.Context, phoneNumber string) (user model.Staff, err error)
}

type StaffRepositoryInfra struct {
	DB *infra.PostgresConn
}

func ProvideStaffRepositoryInfra(db *infra.PostgresConn) *StaffRepositoryInfra {
	r := new(StaffRepositoryInfra)
	r.DB = db
	return r
}
