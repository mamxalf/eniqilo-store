package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/response"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// InsertNewProduct adds a new Product.
// @Summary Add a new Product
// @Description This endpoint is used to add a new Product.
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body request.InsertProductRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/product [post]
func (h *ProductHandler) InsertNewProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var productRequest request.InsertProductRequest
	if err := decoder.Decode(&productRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := productRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	staffID, err := uuid.Parse(claimStaff["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	productRequest.StaffID = staffID
	res, err := h.ProductService.InsertNewProduct(r.Context(), productRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// Find Product by id.
// @Summary GetProduct.
// @Description This endpoint for get Product data by ID.
// @Tags Product
// @Accept  json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID by Product"
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/product/{id} [get]
func (h *ProductHandler) Find(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	staffID, err := uuid.Parse(claimStaff["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}

	res, err := h.ProductService.GetProductData(r.Context(), staffID, idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// FindAllProductData retrieves all product data based on optional filtering parameters.
// @Summary GetAllProduct
// @Description This endpoint retrieves all product data based on optional filtering parameters.
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param id query string false "Limit the output based on the product's ID"
// @Param limit query int false "Limit the number of items in the response (default 5)" default(5)
// @Param offset query int false "Offset the start point of data return (default 0)" default(0)
// @Param name query string false "Filter based on product's name"
// @Param isAvailable query bool false "Filter based on availability" Enums(true, false)
// @Param category query string false "Filter based on category" Enums(Clothing, Accessories, Footwear, Beverages)
// @Param sku query string false "Filter based on product's SKU"
// @Param price query string false "Sort by price" Enums(asc, desc)
// @Param inStock query bool false "Check whether the product is in stock" Enums(true, false)
// @Param owned query bool false "Filter if the user owns the product or not" Enums(true, false)
// @Param search query string false "Display information containing the search term in the name of the product"
// @Param createdAt query string false "Sort by created time" Enums(asc, desc)
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/product/ [get]
func (h *ProductHandler) FindAllProductData(w http.ResponseWriter, r *http.Request) {
	var request request.ProductQueryParams
	query := r.URL.Query()

	// Retrieve values from query parameters
	request.Search = query.Get("search")
	request.ID = query.Get("id")

	// Assign other query parameters
	request.Name = query.Get("name")
	request.Category = query.Get("category")
	request.SKU = query.Get("sku")
	request.Price = query.Get("price")

	// For numerical and boolean values, we need to safely parse them
	// because they require conversion from string
	if limit, err := strconv.Atoi(query.Get("limit")); err == nil {
		request.Limit = limit
	} else {
		request.Limit = 5 // default value if not specified or error in conversion
	}

	if offset, err := strconv.Atoi(query.Get("offset")); err == nil {
		request.Offset = offset
	} else {
		request.Offset = 0 // default value if not specified or error in conversion
	}

	// Parse boolean values
	inStockStr := query.Get("inStock")
	if inStock, err := strconv.ParseBool(inStockStr); err == nil {
		request.InStock = inStock
	}

	isAvailableStr := query.Get("isAvailable")
	if isAvailable, err := strconv.ParseBool(isAvailableStr); err == nil {
		request.IsAvailable = isAvailable
	} else {
		request.IsAvailable = true
	}
	ownedStr := query.Get("owned")
	if owned, err := strconv.ParseBool(ownedStr); err == nil {
		request.Owned = owned
	}

	// Get user ID from JWT claim
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}

	staffID, err := uuid.Parse(claimStaff["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}

	// Call ProductService to get all product data
	res, err := h.ProductService.GetAllProductData(r.Context(), staffID, request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}

// UpdateProductData updates an existing Product.
// @Summary Update a Product
// @Description This endpoint is used to update an existing Product.
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID of the product to be updated"
// @Param request body request.UpdateProductRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/product/{id} [put]
func (h *ProductHandler) UpdateProductData(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(idStr)
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format product id")
		response.WithError(w, err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var productRequest request.UpdateProductRequest
	if err := decoder.Decode(&productRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := productRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// Get user ID from JWT claim
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	_, err = uuid.Parse(claimStaff["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	res, err := h.ProductService.UpdateProductData(r.Context(), productID, productRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// DeleteProductData deletes an existing Product.
// @Summary Delete a Product
// @Description This endpoint is used to delete an existing Product.
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path string true "ID of the product to be deleted"
// @Success 200 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/product/{id} [delete]
func (h *ProductHandler) DeleteProductData(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	_, err := uuid.Parse(claimStaff["ownerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}
	res, err := h.ProductService.DeleteProductData(r.Context(), idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
