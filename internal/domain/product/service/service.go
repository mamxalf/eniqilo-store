package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/repository"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
)

type ProductService interface {
	GetAllProductData(ctx context.Context, params request.ProductFilterParams) (res []model.Product, err error)
}
type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Config            *config.Config
}

// ProvideProductServiceImpl is the provider for this service.
func ProvideProductServiceImpl(
	productRepository repository.ProductRepository,
	config *config.Config,
) *ProductServiceImpl {
	s := new(ProductServiceImpl)
	s.ProductRepository = productRepository
	s.Config = config
	return s
}
