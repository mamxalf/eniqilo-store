package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

func (p *ProductServiceImpl) GetAllProductData(ctx context.Context, params request.ProductFilterParams) (res []model.Product, err error) {
	res, err = p.ProductRepository.FindAll(ctx, params)
	if err != nil {
		logger.ErrorWithMessage(err, "failed to get product data")
		return
	}
	if len(res) == 0 {
		res = []model.Product{}
		return
	}
	return
}
