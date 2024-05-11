package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"fmt"

	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

func (p *ProductRepositoryInfra) InsertTransaction(ctx context.Context, transaction model.InsertTransaction) (err error) {
	query := `INSERT INTO transactions (customer_id, product_details, paid, change)
				  VALUES ($1, $2, $3, $4)
				  RETURNING id, customer_id, product_details, paid, change;`
	productDetailsJSON, err := json.Marshal(transaction.ProductDetails)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	_, err = p.DB.PG.ExecContext(ctx, query, transaction.CustomerID, productDetailsJSON, transaction.Paid, transaction.Change)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

// func (p *ProductRepositoryInfra) FindTransaction(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (transactions []model.Transaction, err error) {
// 	query := `
// 		SELECT id, customer_id, product_details, paid, change, created_at, updated_at
// 		FROM transactions
// 		WHERE customer_id = ?`

// 	var args []interface{}
// 	args = append(args, customerID)

// 	if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
// 		query += fmt.Sprintf(" ORDER BY created_at %s", params.CreatedAt)
// 	} else {
// 		logger.Info("Invalid createdAt value. Ignoring the parameter.")
// 	}

// 	query += " LIMIT ? OFFSET ?"
// 	args = append(args, params.Limit, params.Offset)

// 	err = p.DB.PG.SelectContext(ctx, &transactions, query, args...)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			err = failure.NotFound("Transactions not found")
// 			return nil, err
// 		}
// 		logger.ErrorWithStack(err)
// 		err = failure.InternalError(err)
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (p *ProductRepositoryInfra) FindTransaction(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (transactions []model.Transaction, err error) {
// 	query := `
// 		SELECT id, customer_id, product_details, paid, change, created_at, updated_at
// 		FROM transactions
// 		WHERE customer_id = $1`

// 	var args []interface{}
// 	args = append(args, customerID)

// 	if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
// 		query += fmt.Sprintf(" ORDER BY created_at %s", params.CreatedAt)
// 	} else {
// 		logger.Info("Invalid createdAt value. Ignoring the parameter.")
// 	}

// 	query += " LIMIT $2 OFFSET $3"
// 	args = append(args, params.Limit, params.Offset)

// 	err = p.DB.PG.SelectContext(ctx, &transactions, query, args...)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			err = failure.NotFound("Transactions not found")
// 			return nil, err
// 		}
// 		logger.ErrorWithStack(err)
// 		err = failure.InternalError(err)
// 		return nil, err
// 	}

// 	return transactions, nil
// }

func (p *ProductRepositoryInfra) FindTransaction(ctx context.Context, customerID uuid.UUID, params request.TransactionQueryParams) (transactions []model.Transaction, err error) {
	query := `
		SELECT id, customer_id, product_details, paid, change, created_at, updated_at
		FROM transactions
		WHERE customer_id = $1`

	var args []interface{}
	args = append(args, customerID)

	if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
		query += fmt.Sprintf(" ORDER BY created_at %s", params.CreatedAt)
	} else {
		logger.Info("Invalid createdAt value. Ignoring the parameter.")
	}

	query += " LIMIT $2 OFFSET $3"
	args = append(args, params.Limit, params.Offset)

	err = p.DB.PG.SelectContext(ctx, &transactions, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Transactions not found")
			return nil, err
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}

	return transactions, nil
}
