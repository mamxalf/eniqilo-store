package repository

import (
	"context"

	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"

	"github.com/google/uuid"
)

type ProductRepository interface {
	// Insert Product CRUD Interface
	Insert(ctx context.Context, product model.InsertProduct) (newProduct *model.Product, err error)
	Find(ctx context.Context, productID uuid.UUID) (product model.Product, err error)
	FindAll(ctx context.Context, staffID uuid.UUID, params request.ProductQueryParams) (products []model.Product, err error)
	Update(ctx context.Context, productID uuid.UUID, product model.Product) (updatedProduct *model.Product, err error)
	Delete(ctx context.Context, productID uuid.UUID) (deletedID uuid.UUID, err error)
}

type ProductRepositoryInfra struct {
	DB *infra.PostgresConn
}

func ProvideProductRepositoryInfra(db *infra.PostgresConn) *ProductRepositoryInfra {
	r := new(ProductRepositoryInfra)
	r.DB = db
	return r
}
