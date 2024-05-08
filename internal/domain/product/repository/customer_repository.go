package repository

import (
	"context"
	"fmt"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"strings"
)

func (r *ProductRepositoryInfra) FindAll(ctx context.Context, params request.ProductFilterParams) (cats []model.Product, err error) {
	var products []model.Product
	var conditions []string
	var args []interface{}

	baseQuery := "SELECT * FROM products WHERE 1=1"

	if params.Name != "" {
		conditions = append(conditions, "name ILIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+params.Name+"%")
	}

	if params.Category != "" {
		conditions = append(conditions, "category = $"+fmt.Sprint(len(args)+1))
		args = append(args, params.Category)
	}

	if params.SKU != "" {
		conditions = append(conditions, "sku = $"+fmt.Sprint(len(args)+1))
		args = append(args, params.SKU)
	}

	if params.InStock != nil {
		conditions = append(conditions, "in_stock = $"+fmt.Sprint(len(args)+1))
		args = append(args, *params.InStock)
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if params.Price != "" {
		order := "ASC"
		if params.Price == "desc" {
			order = "DESC"
		}
		baseQuery += fmt.Sprintf(" ORDER BY price %s", order)
	}

	if params.Limit > 0 {
		baseQuery += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		args = append(args, params.Limit)
	}

	if params.Offset > 0 {
		baseQuery += fmt.Sprintf(" OFFSET $%d", len(args)+1)
		args = append(args, params.Offset)
	}

	err = r.DB.PG.SelectContext(ctx, &products, baseQuery, args...)
	if err != nil {
		return
	}

	return
}
