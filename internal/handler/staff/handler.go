package staff

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/service"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/response"
	"net/http"
)

type StaffHandler struct {
	StaffService  service.StaffService
	JWTMiddleware *middleware.JWT
}

func ProvideStaffHandler(staffService service.StaffService, jwtMiddleware *middleware.JWT) StaffHandler {
	return StaffHandler{
		StaffService:  staffService,
		JWTMiddleware: jwtMiddleware,
	}
}

func (h *StaffHandler) Router(r chi.Router) {
	r.Route("/staff", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.Register)
			r.Post("/login", h.Login)
		})
	})
}

// Register sign up staff.
// @Summary Register Staff
// @Description This endpoint for Register Staff.
// @Tags Staff
// @Accept  json
// @Produce json
// @Param request body request.RegisterRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/staff/register [post]
func (h *StaffHandler) Register(w http.ResponseWriter, r *http.Request) {
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
	res, err := h.StaffService.RegisterNewStaff(r.Context(), registerRequest)
	if err != nil {
		logger.WarningWithMessage(err, "error registering new staff")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// Login sign in staff.
// @Summary Login Staff
// @Description This endpoint for Login Staff.
// @Tags Staff
// @Accept  json
// @Produce json
// @Param request body request.LoginRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/staff/login [post]
func (h *StaffHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginRequest request.LoginRequest
	if err := decoder.Decode(&loginRequest); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := loginRequest.Validate(); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.StaffService.LoginStaff(r.Context(), loginRequest)
	if err != nil {
		logger.WarningWithMessage(err, "error login staff")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, res)
}
