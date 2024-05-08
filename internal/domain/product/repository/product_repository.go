package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mamxalf/eniqilo-store/internal/domain/product/model"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"

	"github.com/google/uuid"
)

var productQueries = struct {
	InsertProduct string
	GetProduct    string
	GetAllProduct string
	DeleteProduct string
}{
	InsertProduct: "INSERT INTO products %s VALUES %s RETURNING id",
	GetProduct:    "SELECT * FROM products WHERE 1=1",
	GetAllProduct: "SELECT * FROM products %s",
	DeleteProduct: "DELETE FROM products WHERE id = $1",
}

func (p *ProductRepositoryInfra) Insert(ctx context.Context, product model.InsertProduct) (newProduct *model.Product, err error) {
	query := `INSERT INTO products (staff_id, name, sku, category, imageurl, notes, price, stock, location, isavailable)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
              RETURNING id, staff_id, name, sku, category, imageurl, notes, price, stock, location, isavailable;`
	newProduct = &model.Product{}
	err = p.DB.PG.QueryRowxContext(ctx, query, product.StaffID, product.Name, product.SKU, product.Category, product.ImageURL, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable).StructScan(newProduct)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}
	return newProduct, nil
}

func (p *ProductRepositoryInfra) Find(ctx context.Context, productID uuid.UUID) (product model.Product, err error) {
	whereClauses := " WHERE id = $1 LIMIT 1"
	query := fmt.Sprintf(productQueries.GetAllProduct, whereClauses)
	err = p.DB.PG.GetContext(ctx, &product, query, productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Product not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

func (p *ProductRepositoryInfra) FindAll(ctx context.Context, StaffID uuid.UUID, params request.ProductQueryParams) (products []model.Product, err error) {
	baseQuery := productQueries.GetAllProduct // Assuming this starts with "SELECT ... FROM products WHERE 1=1"
	var args []interface{}
	var conditions []string

	// Always include the staff_id in the query
	if params.Owned {
		conditions = append(conditions, fmt.Sprintf("staff_id = $%d", len(args)+1))
		args = append(args, StaffID)
	}

	// Check if ID is specified
	if params.ID != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(args)+1))
		args = append(args, params.ID)
	}

	// Check if SKU is specified
	if params.SKU != "" {
		conditions = append(conditions, fmt.Sprintf("sku = $%d", len(args)+1))
		args = append(args, params.SKU)
	}

	// Check if Category is specified
	if params.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(args)+1))
		args = append(args, params.Category)
	}

	// Check if Search term is specified
	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)+1))
		args = append(args, "%"+params.Search+"%")
	}

	// Check if product is available
	if params.IsAvailable {
		conditions = append(conditions, "isavailable = true")
	}

	// Adding the conditions to the base query
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Adding sorting by price
	if params.Price != "" {
		switch params.Price {
		case "asc":
			baseQuery += " ORDER BY price ASC"
		case "desc":
			baseQuery += " ORDER BY price DESC"
		}
	}

	// Adding sorting by created time
	if params.CreatedAt != "" {
		switch params.CreatedAt {
		case "asc":
			baseQuery += " ORDER BY created_at ASC"
		case "desc":
			baseQuery += " ORDER BY created_at DESC"
		}
	}

	// Adding pagination with proper indexing
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, params.Limit, params.Offset)

	// Executing the query
	err = p.DB.PG.SelectContext(ctx, &products, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryInfra) Update(ctx context.Context, productID uuid.UUID, product model.Product) (updatedProduct *model.Product, err error) {
	var setParts []string
	var args []interface{}
	argID := 1

	// Dynamically build the SQL query based on provided fields
	if product.Name != "" {
		setParts = append(setParts, "name = $"+strconv.Itoa(argID))
		args = append(args, product.Name)
		argID++
	}
	if product.SKU != "" {
		setParts = append(setParts, "sku = $"+strconv.Itoa(argID))
		args = append(args, product.SKU)
		argID++
	}
	if product.Category != "" {
		setParts = append(setParts, "category = $"+strconv.Itoa(argID))
		args = append(args, product.Category)
		argID++
	}
	if product.ImageURL != "" {
		setParts = append(setParts, "imageurl = $"+strconv.Itoa(argID))
		args = append(args, product.ImageURL)
		argID++
	}
	if product.Notes != "" {
		setParts = append(setParts, "notes = $"+strconv.Itoa(argID))
		args = append(args, product.Notes)
		argID++
	}
	if product.Price != 0 {
		setParts = append(setParts, "price = $"+strconv.Itoa(argID))
		args = append(args, product.Price)
		argID++
	}
	if product.Stock != 0 {
		setParts = append(setParts, "stock = $"+strconv.Itoa(argID))
		args = append(args, product.Stock)
		argID++
	}
	if product.Location != "" {
		setParts = append(setParts, "location = $"+strconv.Itoa(argID))
		args = append(args, product.Location)
		argID++
	}

	if len(setParts) == 0 {
		return // No updates to perform
	}

	// Construct the full SQL statement
	updateQuery := "UPDATE products SET " + strings.Join(setParts, ", ") + " WHERE id = $" + strconv.Itoa(argID) + " RETURNING *"
	args = append(args, productID)

	// Execute the query
	updatedProduct = &model.Product{}
	err = p.DB.PG.GetContext(ctx, updatedProduct, updateQuery, args...)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}
	return
}
func (p *ProductRepositoryInfra) Delete(ctx context.Context, productID uuid.UUID) (deletedID uuid.UUID, err error) {
	result, err := p.DB.PG.ExecContext(ctx, productQueries.DeleteProduct, productID)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	if rowsAffected == 0 {
		err = failure.NotFound("Product not found!")
		return
	}

	deletedID = productID
	return
}
