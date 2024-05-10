package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

var transactionQueries = struct {
	InsertTransaction string
	GetTransaction    string
}{
	InsertTransaction: "INSERT INTO transactions %s VALUES %s RETURNING id",
	GetTransaction:    "SELECT * FROM transactions %s",
}

func (p *ProductRepositoryInfra) InsertTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	query := `INSERT INTO transactions (customer_id, product_details, paid, change)
				  VALUES ($1, $2, $3, $4, $5)
				  RETURNING id, customer_id, product_details, paid, change;`
	newTransaction := &model.Transaction{}
	productDetailsJSON, err := json.Marshal(transaction.ProductDetails)
	if err != nil {
		logger.ErrorWithStack(err)
		return nil, failure.InternalError(err)
	}
	err = p.DB.PG.QueryRowxContext(ctx, query, transaction.CustomerID, productDetailsJSON, transaction.Paid, transaction.Change).StructScan(newTransaction)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}
	return newTransaction, nil
}

func (p *ProductRepositoryInfra) FindTransaction(ctx context.Context, customerID uuid.UUID) (transaction model.Transaction, err error) {
	whereClauses := " WHERE id = $1 LIMIT 1"
	query := fmt.Sprintf(transactionQueries.GetTransaction, whereClauses)
	err = p.DB.PG.GetContext(ctx, &transaction, query, customerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Transaction not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}
