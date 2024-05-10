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
