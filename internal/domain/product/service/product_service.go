package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"

	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *ProductServiceImpl) InsertNewProduct(ctx context.Context, req request.InsertProductRequest) (res response.ProductResponse, err error) {
	product, err := req.ToModel()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}

	result, err := u.ProductRepository.Insert(ctx, product)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("can't insert new product")
		return
	}

	res = response.ProductResponse{
		ID:          result.ID,
		Name:        result.Name,
		SKU:         result.SKU,
		Category:    result.Category,
		ImageURL:    result.ImageURL,
		Notes:       result.Notes,
		Price:       result.Price,
		Stock:       result.Stock,
		Location:    result.Location,
		IsAvailable: result.IsAvailable,
	}
	return
}
func (u *ProductServiceImpl) GetProductData(ctx context.Context, staffID uuid.UUID, productID string) (res response.ProductResponse, err error) {
	id, err := uuid.Parse(productID)
	if err != nil {
		return
	}
	product, err := u.ProductRepository.Find(ctx, id)
	if err != nil {
		return
	}

	res = response.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		SKU:         product.SKU,
		Category:    product.Category,
		ImageURL:    product.ImageURL,
		Notes:       product.Notes,
		Price:       product.Price,
		Stock:       product.Stock,
		Location:    product.Location,
		IsAvailable: product.IsAvailable,
	}
	return
}

func (u *ProductServiceImpl) GetAllProductData(ctx context.Context, staffID uuid.UUID, params request.ProductQueryParams) (res []model.Product, err error) {
	res, err = u.ProductRepository.FindAll(ctx, staffID, params)
	if err != nil {
		return
	}
	if len(res) == 0 {
		return []model.Product{}, nil
	}

	return
}
func (u *ProductServiceImpl) UpdateProductData(ctx context.Context, productID uuid.UUID, req request.UpdateProductRequest) (res response.ProductResponse, err error) {
	productModel, err := req.ToModel()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}

	product, err := u.ProductRepository.Update(ctx, productID, productModel)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	res = response.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		SKU:         product.SKU,
		Category:    product.Category,
		ImageURL:    product.ImageURL,
		Notes:       product.Notes,
		Price:       product.Price,
		Stock:       product.Stock,
		Location:    product.Location,
		IsAvailable: product.IsAvailable,
	}
	return
}

func (u *ProductServiceImpl) DeleteProductData(ctx context.Context, productIDStr string) (res response.ProductResponse, err error) {
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		return
	}

	_, err = u.ProductRepository.Delete(ctx, productID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProductData - Service] Internal Error")
		return
	}

	return
}
