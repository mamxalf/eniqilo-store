package product

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/response"

	"github.com/google/uuid"
)

// InsertNewTransaction adds a new Transaction.
// @Summary Add a new Transaction
// @Description This endpoint is used to add a new Transaction.
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body request.InsertTransactionRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/product/checkout [post]
func (h *ProductHandler) InsertNewTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var transactionRequest request.InsertTransactionRequest
	if err := decoder.Decode(&transactionRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := transactionRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err := h.ProductService.InsertNewTransaction(r.Context(), transactionRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusOK, "Succes Insert request")
}

// GetTransactionHistory retrieves transaction history for a customer based on optional filtering parameters.
// @Summary Get Transaction History
// @Description Retrieve transaction history for a customer based on optional filtering parameters.
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerToken
// @Param customerId query string true "Customer ID"
// @Param limit query int false "Limit the number of items in the response (default 5)" default(5)
// @Param offset query int false "Offset the start point of data return (default 0)" default(0)
// @Param createdAt query string false "Sort by created time" Enums(asc, desc)
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/product/checkout/history [get]
func (h *ProductHandler) FindTransaction(w http.ResponseWriter, r *http.Request) {
	var params request.TransactionQueryParams
	// fmt.Println("res", params)
	query := r.URL.Query()

	// Retrieve values from query parameters
	params.CustomerID = query.Get("customerId")

	// For numerical values, we need to safely parse them
	// because they require conversion from string
	if limit, err := strconv.Atoi(query.Get("limit")); err == nil {
		params.Limit = limit
	} else {
		params.Limit = 5 // default value if not specified or error in conversion
	}

	if offset, err := strconv.Atoi(query.Get("offset")); err == nil {
		params.Offset = offset
	} else {
		params.Offset = 0 // default value if not specified or error in conversion
	}

	// Parse createdAt parameter
	params.CreatedAt = query.Get("createdAt")

	// Parse customer ID from params
	customerID, err := uuid.Parse(params.CustomerID)
	if err != nil {
		logger.ErrorWithMessage(err, "Invalid format for customer ID")
		err = failure.BadRequest(err)
		response.WithError(w, err)
		return
	}
	// Call ProductService to get transaction history
	res, err := h.ProductService.GetTransactionData(r.Context(), customerID, params)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
