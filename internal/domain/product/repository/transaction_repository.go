package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	// "fmt"

	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
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

func (p *ProductRepositoryInfra) FindTransaction(ctx context.Context, customerID uuid.UUID) (transaction model.Transaction, err error) {
	query := `
		SELECT transaction_id, customer_id, product_details, paid, change, created_at
		FROM transactions
		WHERE customer_id = $1
		ORDER BY created_at DESC
		LIMIT 5 OFFSET 0
	`

	err = p.DB.PG.GetContext(ctx, &transaction, query, customerID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Transaction not found")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return transaction, nil
}
