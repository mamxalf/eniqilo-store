package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

func (u *ProductServiceImpl) InsertNewTransaction(ctx context.Context, req request.InsertTransactionRequest) (res response.TransactionResponse, err error) {

	transaction, err := req.ToModel()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}

	result, err := u.ProductRepository.InsertTransaction(ctx, transaction)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("can't insert new transaction")
		return
	}
	productDetails := make([]response.ProductDetail, len(result.ProductDetails))
	for i, detail := range result.ProductDetails {
		productDetails[i] = response.ProductDetail{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
		}
	}
	res = response.TransactionResponse{
		TransactionID:  result.ID,
		CustomerID:     result.CustomerID,
		ProductDetails: productDetails,
		Paid:           result.Paid,
		Change:         result.Change,
		CreatedAt:      result.CreatedAt,
	}
	return
}

func (u *ProductServiceImpl) GetTransactionData(ctx context.Context, customerID uuid.UUID, transactionID string) (res response.TransactionResponse, err error) {
	id, err := uuid.Parse(transactionID)
	if err != nil {
		return res, err
	}

	transaction, err := u.ProductRepository.FindTransaction(ctx, id)
	if err != nil {
		return res, err
	}

	productDetails := make([]response.ProductDetail, len(transaction.ProductDetails))
	for i, detail := range transaction.ProductDetails {
		productDetails[i] = response.ProductDetail{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
		}
	}

	res = response.TransactionResponse{
		TransactionID:  transaction.ID,
		CustomerID:     transaction.CustomerID,
		ProductDetails: productDetails,
		Paid:           transaction.Paid,
		Change:         transaction.Change,
	}
	return res, nil
}
