package service

import (
	"context"

	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/repository"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/response"

	"github.com/google/uuid"
)

type ProductService interface {
	// Product Service Interface
	InsertNewProduct(ctx context.Context, req request.InsertProductRequest) (res response.ProductResponse, err error)
	GetProductData(ctx context.Context, staffID uuid.UUID, productID string) (res response.ProductResponse, err error)
	GetAllProductData(ctx context.Context, staffID uuid.UUID, params request.ProductQueryParams) (res []response.ProductResponse, err error)
	UpdateProductData(ctx context.Context, productID uuid.UUID, req request.UpdateProductRequest) (res response.ProductResponse, err error)
	DeleteProductData(ctx context.Context, productID string) (res response.ProductResponse, err error)

	// product customer sku
	GetAllProductCustomerData(ctx context.Context, params request.ProductSKUFilterParams) (res []model.ProductSKU, err error)

	//transaction customer
	InsertNewTransaction(ctx context.Context, req request.InsertTransactionRequest) (err error)
	GetTransactionData(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (res []response.TransactionResponse, err error)
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
