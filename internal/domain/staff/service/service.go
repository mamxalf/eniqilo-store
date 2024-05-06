package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/repository"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/response"
)

type StaffService interface {
	RegisterNewStaff(ctx context.Context, req request.RegisterRequest) (res response.SessionResponse, err error)
}

type StaffServiceImpl struct {
	StaffRepository repository.StaffRepository
	Config          *config.Config
}

// ProvideStaffServiceImpl is the provider for this service.
func ProvideStaffServiceImpl(
	userRepository repository.StaffRepository,
	config *config.Config,
) *StaffServiceImpl {
	s := new(StaffServiceImpl)
	s.StaffRepository = userRepository
	s.Config = config
	return s
}
