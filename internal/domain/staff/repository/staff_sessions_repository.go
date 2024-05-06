package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/model"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

type StaffRegister struct {
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
}

func (u *StaffRepositoryInfra) Register(ctx context.Context, staffRegister *model.StaffRegister) (id uuid.UUID, err error) {
	stmt, err := u.DB.PG.PrepareContext(ctx, "INSERT INTO staffs (name, phone, password) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	// Extract fields from staffRegister and pass them to QueryRowContext
	err = stmt.QueryRowContext(ctx, staffRegister.Name, staffRegister.Phone, staffRegister.Password).Scan(&id)
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
