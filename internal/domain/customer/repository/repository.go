package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/model"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

type CustomerRepository interface {
	Register(ctx context.Context, customerRegister *model.CustomerRegister) (id uuid.UUID, err error)
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
	stmt, err := r.DB.PG.PrepareContext(ctx, "INSERT INTO customers (name, phone) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	// Extract fields from staffRegister and pass them to QueryRowContext
	err = stmt.QueryRowContext(ctx, customerRegister.Name, customerRegister.Phone).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// Check if the error code is for a unique violation
			if pqErr.Code == "23505" {
				err = failure.Conflict("Duplicate key error occurred", pqErr.Message)
				return
			}
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}
