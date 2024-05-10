package customer

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/customer/service"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/response"
	"net/http"
)

type CustomerHandler struct {
	CustomerService service.CustomerService
	JWTMiddleware   *middleware.JWT
}

func ProvideCustomerHandler(customerService service.CustomerService, jwtMiddleware *middleware.JWT) CustomerHandler {
	return CustomerHandler{
		CustomerService: customerService,
		JWTMiddleware:   jwtMiddleware,
	}
}

func (h *CustomerHandler) Router(r chi.Router) {
	r.Route("/customer", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(h.JWTMiddleware.VerifyToken)
			r.Post("/", h.GetCustomers)
		})
		r.Group(func(r chi.Router) {
			r.Post("/register", h.Register)
		})
	})
}

// Register customer from cashier.
// @Summary Register Customer
// @Description This endpoint for Register customer from cashier.
// @Tags Customer
// @Accept  json
// @Produce json
// @Param request body request.RegisterRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/customer/register [post]
func (h *CustomerHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var registerRequest request.RegisterRequest
	if err := decoder.Decode(&registerRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := registerRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.CustomerService.RegisterNewCustomer(r.Context(), registerRequest)
	if err != nil {
		logger.WarningWithMessage(err, "error registering new customer")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// GetCustomers godoc
// @Summary Get customers
// @Description Retrieves customers based on provided search criteria.
// @Tags customers
// @Accept json
// @Produce json
// @Param phoneNumber query string false "Search customer by phone number without the '+' sign, uses wildcard matching (ex: '62' matches '+628123...')"
// @Param name query string false "Search customer by name, uses wildcard and case insensitive matching (ex: 'een' matches 'kayleena')"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /customer [get]
func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := &request.CustomerQueryParams{
		PhoneNumber: query.Get("phoneNumber"),
		Name:        query.Get("name"),
	}

	if err := params.Validate(); err != nil {
		logger.ErrorWithMessage(err, "Invalid parameters")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	res, err := h.CustomerService.GetAllRegisteredCustomers(r.Context(), *params)
	if err != nil {
		logger.WarningWithMessage(err, "Error getting customers data")
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
