package service

import (
	"context"

	"github.com/mamxalf/eniqilo-store/config"
	productRepository "github.com/mamxalf/eniqilo-store/internal/domain/product/repository"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/response"
	staffRepository "github.com/mamxalf/eniqilo-store/internal/domain/staff/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	// Product Service Interface
	InsertNewProduct(ctx context.Context, req request.InsertProductRequest) (res response.ProductResponse, err error)
	GetProductData(ctx context.Context, staffID uuid.UUID, productID string) (res response.ProductResponse, err error)
	GetAllProductData(ctx context.Context, staffID uuid.UUID, params request.ProductQueryParams) (res []response.ProductResponse, err error)
	UpdateProductData(ctx context.Context, productID uuid.UUID, req request.UpdateProductRequest) (res response.ProductResponse, err error)
	DeleteProductData(ctx context.Context, productID string) (res response.ProductResponse, err error)
}

type ProductServiceImpl struct {
	ProductRepository productRepository.ProductRepository
	staffRepository   staffRepository.StaffRepository
	Config            *config.Config
}

// ProvideProductServiceImpl is the provider for this service.
func ProvideProductServiceImpl(
	productRepository productRepository.ProductRepository,
	staffRepository staffRepository.StaffRepository,
	config *config.Config,
) *ProductServiceImpl {
	s := new(ProductServiceImpl)
	s.ProductRepository = productRepository
	s.staffRepository = staffRepository
	s.Config = config
	return s
}
