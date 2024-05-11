package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strconv"
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/request"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/response"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
	claimCustomer, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	customerID, err := uuid.Parse(claimCustomer["customerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format customer_id")
		response.WithError(w, err)
		return
	}
	transactionRequest.CustomerID = customerID
	fmt.Println("transaction request:", transactionRequest)
	res, err := h.ProductService.InsertNewTransaction(r.Context(), transactionRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// GetCheckoutHistory handles the retrieval of checkout history.
// @Summary Get Checkout History
// @Description This endpoint retrieves checkout history for a customer.
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerToken
// @Param customer_id query string false "Customer ID"
// @Param limit query int false "Limit the number of transactions (default: 5)"
// @Param offset query int false "Offset the list of transactions (default: 0)"
// @Param sort query string false "Sort order for transactions (asc: oldest first, desc: newest first)"
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/product/checkout/history [get]
func (h *ProductHandler) FindTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	idStr := queryParams.Get("transaction_id")
	claimStaff, ok := middleware.GetClaimsUser(r.Context()).(jwt.MapClaims)
	if !ok {
		log.Warn().Msg("invalid claim jwt")
		err := failure.Unauthorized("invalid claim jwt")
		response.WithError(w, err)
		return
	}
	customerID, err := uuid.Parse(claimStaff["customerID"].(string))
	if err != nil {
		log.Warn().Msg(err.Error())
		err = failure.Unauthorized("invalid format user_id")
		response.WithError(w, err)
		return
	}

	res, err := h.ProductService.GetTransactionData(r.Context(), customerID, idStr)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// GetTransactionData handles the retrieval of transaction data for a customer.
// @Summary Get Transaction Data
// @Description This endpoint retrieves transaction data for a customer.
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerToken
// @Param customerId query string false "Customer ID" Format: UUID
// @Success 200 {object} response.TransactionResponse "Successful operation"
// @Failure 400 {object} response.Base "Bad request"
// @Failure 401 {object} response.Base "Unauthorized"
// @Failure 500 {object} response.Base "Internal server error"
// @Router /v1/product/checkout/history [get]
func (h *ProductHandler) GetTransactionData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	customerIDStr := queryParams.Get("customerId")
	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		log.Warn().Msg("Invalid customer ID format")
		err := failure.BadRequest(err)
		response.WithError(w, err)
		return
	}

	// Call service to get transaction data
	res, err := h.ProductService.GetTransactionData(r.Context(), customerID, "")
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
