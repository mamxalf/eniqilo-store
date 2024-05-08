package repository

import (
	"context"
	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
)

type ProductRepository interface {
	FindAll(ctx context.Context, params request.ProductFilterParams) (cats []model.Product, err error)
}

type ProductRepositoryInfra struct {
	DB *infra.PostgresConn
}

func ProvideProductRepositoryInfra(db *infra.PostgresConn) *ProductRepositoryInfra {
	r := new(ProductRepositoryInfra)
	r.DB = db
	return r
}
