package product

import (
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/response"
	"net/http"
	"strconv"
)

// SearchSKUProduct godoc
// @Summary Get products with filters
// @Description Retrieves products based on provided filtering criteria.
// @Tags Products
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of items in the response (default 5)" default(5)
// @Param offset query int false "Offset the start point of data return (default 0)" default(0)
// @Param name query string false "Filter by product name using a case insensitive wildcard match"
// @Param category query string false "Filter by category" Enums(Clothing, Accessories, Footwear, Beverages)
// @Param sku query string false "Filter by product SKU"
// @Param price query string false "Sort by price" Enums(asc, desc)
// @Param inStock query bool false "Filter products that are in stock"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /product/customer [get]
func (h *ProductHandler) SearchSKUProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := &request.ProductFilterParams{
		Limit:  5, // Default value
		Offset: 0, // Default value
	}

	// Manually parse query parameters
	if limitStr := query.Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			logger.ErrorWithMessage(err, "Invalid limit parameter")
			response.WithError(w, failure.BadRequest(err))
			return
		}
		params.Limit = limit
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			logger.ErrorWithMessage(err, "Invalid offset parameter")
			response.WithError(w, failure.BadRequest(err))
			return
		}
		params.Offset = offset
	}

	params.Name = query.Get("name")
	params.Category = query.Get("category")
	params.SKU = query.Get("sku")
	params.Price = query.Get("price")

	if inStockStr := query.Get("inStock"); inStockStr != "" {
		inStock, err := strconv.ParseBool(inStockStr)
		if err != nil {
			logger.ErrorWithMessage(err, "Invalid inStock parameter")
			response.WithError(w, failure.BadRequest(err))
			return
		}
		params.InStock = &inStock
	}
	if err := params.Validate(); err != nil {
		logger.ErrorWithMessage(err, "Invalid parameters")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	res, err := h.ProductService.GetAllProductData(r.Context(), *params)
	if err != nil {
		logger.WarningWithMessage(err, "Error getting product data")
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
