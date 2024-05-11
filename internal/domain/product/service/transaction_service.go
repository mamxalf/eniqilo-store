package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
)

func (u *ProductServiceImpl) InsertNewTransaction(ctx context.Context, req request.InsertTransactionRequest) (err error) {

	transaction, err := req.ToModel()

	if err != nil {
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}

	err = u.ProductRepository.InsertTransaction(ctx, transaction)
	if err != nil {
		err = failure.BadRequestFromString("can't insert new transaction")
		return
	}
	return
}

// func (u *ProductServiceImpl) GetTransactionData(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (res []response.TransactionResponse, err error) {
// 	var transactionList []response.TransactionResponse
// 	transactions, err := u.ProductRepository.FindTransaction(ctx, customerID, params)
// 	if err != nil {
// 		return
// 	}
// 	if len(transactions) == 0 {
// 		return []response.TransactionResponse{}, nil
// 	}
// 	for _, transaction := range transactions {
// 		transactionList = append(transactionList, response.TransactionResponse{
// 			TransactionID:  transaction.ID,
// 			CustomerID:     transaction.CustomerID,
// 			ProductDetails: transaction.ProductDetails,
// 			Paid:           transaction.Paid,
// 			Change:         transaction.Change,
// 			CreatedAt:      transaction.CreatedAt,
// 		})
// 	}
// 	res = transactionList
// 	return
// }

// func (u *ProductServiceImpl) GetTransactionData(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (res []response.TransactionResponse, err error) {
// 	transactions, err := u.ProductRepository.FindTransaction(ctx, customerID, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, transaction := range transactions {
// 		var productDetails []response.ProductDetail
// 		for _, detail := range transaction.ProductDetails {
// 			productDetails = append(productDetails, response.ProductDetail{
// 				ProductID: detail.ProductID,
// 				Quantity:  detail.Quantity,
// 			})
// 		}

// 		res = append(res, response.TransactionResponse{
// 			TransactionID:  transaction.ID,
// 			CustomerID:     transaction.CustomerID,
// 			ProductDetails: productDetails,
// 			Paid:           transaction.Paid,
// 			Change:         transaction.Change,
// 			CreatedAt:      transaction.CreatedAt,
// 			UpdatedAt:      transaction.UpdatedAt,
// 		})
// 	}

// 	return res, nil
// }

func (u *ProductServiceImpl) GetTransactionData(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (res []response.TransactionResponse, err error) {
	transactions, err := u.ProductRepository.FindTransaction(ctx, customerID, params)
	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		var productDetails []response.ProductDetail
		for _, detail := range transaction.ProductDetails {
			productDetails = append(productDetails, response.ProductDetail{
				ProductID: detail.ProductID,
				Quantity:  detail.Quantity,
			})
		}

		res = append(res, response.TransactionResponse{
			TransactionID:  transaction.ID,
			CustomerID:     transaction.CustomerID,
			ProductDetails: productDetails,
			Paid:           transaction.Paid,
			Change:         transaction.Change,
			CreatedAt:      transaction.CreatedAt,
			UpdatedAt:      transaction.UpdatedAt,
		})
	}

	return res, nil
}
