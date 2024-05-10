package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

func (p *ProductServiceImpl) GetAllProductCustomerData(ctx context.Context, params request.ProductSKUFilterParams) (res []model.ProductSKU, err error) {
	res, err = p.ProductRepository.FindAllSKU(ctx, params)
	if err != nil {
		logger.ErrorWithMessage(err, "failed to get product data")
		return
	}
	if len(res) == 0 {
		res = []model.ProductSKU{}
		return
	}
	return
}
