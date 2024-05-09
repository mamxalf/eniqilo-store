package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/repository"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"net/http"
)

type CustomerService interface {
	RegisterNewCustomer(ctx context.Context, req request.RegisterRequest) (res response.CustomerCreatedResponse, err error)
}

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Config             *config.Config
}

// ProvideCustomerServiceImpl is the provider for this service.
func ProvideCustomerServiceImpl(
	userRepository repository.CustomerRepository,
	config *config.Config,
) *CustomerServiceImpl {
	s := new(CustomerServiceImpl)
	s.CustomerRepository = userRepository
	s.Config = config
	return s
}

func (c *CustomerServiceImpl) RegisterNewCustomer(ctx context.Context, req request.RegisterRequest) (res response.CustomerCreatedResponse, err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		logger.ErrorInterfaceWithMessage(err, "failed parse request to model", "params", req)
		return
	}

	id, err := c.CustomerRepository.Register(ctx, &registerModel)
	if err != nil {
		if failure.GetCode(err) == http.StatusConflict {
			logger.ErrorInterfaceWithMessage(err, "duplicate phone", "params", req)
			return
		}
		logger.ErrorInterfaceWithMessage(err, "failed register staff", "params", req)
		return
	}

	res = response.CustomerCreatedResponse{
		UserID:      id.String(),
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	return
}
