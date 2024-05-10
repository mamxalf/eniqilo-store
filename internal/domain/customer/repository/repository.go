package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"strings"
)

type CustomerRepository interface {
	Register(ctx context.Context, customerRegister *model.CustomerRegister) (id uuid.UUID, err error)
	FindAll(ctx context.Context, req request.CustomerQueryParams) (customers []model.Customer, err error)
}

type CustomerRepositoryInfra struct {
	DB *infra.PostgresConn
}

func ProvideCustomerRepositoryInfra(db *infra.PostgresConn) *CustomerRepositoryInfra {
	r := new(CustomerRepositoryInfra)
	r.DB = db
	return r
}

func (r *CustomerRepositoryInfra) Register(ctx context.Context, customerRegister *model.CustomerRegister) (id uuid.UUID, err error) {
	stmt, err := r.DB.PG.PrepareContext(ctx, "INSERT INTO customers (name, phone) VALUES ($1, $2) RETURNING id")
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	// Extract fields from staffRegister and pass them to QueryRowContext
	err = stmt.QueryRowContext(ctx, customerRegister.Name, customerRegister.Phone).Scan(&id)
	if err != nil {
		var pqErr pgx.PgError
		if errors.As(err, &pqErr) {
			// Check if the error code is for a unique violation
			if pqErr.Code == "23505" {
				err = failure.Conflict("Duplicate key error occurred", pqErr.Message)
				return
			}
			logger.ErrorWithStack(err)
			err = failure.InternalError(err)
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

func (r *CustomerRepositoryInfra) FindAll(ctx context.Context, req request.CustomerQueryParams) (customers []model.Customer, err error) {
	var whereClauses []string
	var args []interface{}

	// Start with a base query
	baseQuery := "SELECT id, name, phone_number FROM customers WHERE 1=1"

	// Append conditions based on non-empty parameters
	if req.PhoneNumber != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("phone_number LIKE $%d", len(args)+1))
		args = append(args, "%"+req.PhoneNumber+"%") // Assumes phone numbers are stored with '+'
	}
	if req.Name != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("LOWER(name) LIKE LOWER($%d)", len(args)+1))
		args = append(args, "%"+req.Name+"%")
	}

	// Combine the clauses with the base query
	if len(whereClauses) > 0 {
		baseQuery += " AND " + strings.Join(whereClauses, " AND ")
	}

	// Execute the query
	err = r.DB.PG.SelectContext(ctx, &customers, baseQuery, args...)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	return
}
