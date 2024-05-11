package product

import (
	"encoding/json"
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

	err := h.ProductService.InsertNewTransaction(r.Context(), transactionRequest)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusOK, "Succes Insert request")
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
