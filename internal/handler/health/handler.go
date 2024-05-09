package health

import (
	"github.com/go-chi/chi"
	"github.com/mamxalf/eniqilo-store/infra"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/response"
	"net/http"
)

type HealthHandler struct {
	DB *infra.PostgresConn
}

func ProvideHealthHandler(db *infra.PostgresConn) HealthHandler {
	return HealthHandler{
		DB: db,
	}
}

func (h *HealthHandler) Router(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/ping", h.Ping)
			r.Get("/ping-db", h.PingDB)
		})
	})
}

// Ping get health message.
// @Summary Ping.
// @Description This endpoint for get health message.
// @Tags Health
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/health/ping [get]
func (h *HealthHandler) Ping(w http.ResponseWriter, _ *http.Request) {
	response.WithMessage(w, http.StatusOK, "PONG!!!")
}

// PingDB get health message.
// @Summary PingDB.
// @Description This endpoint for get health message.
// @Tags Health
// @Accept  json
// @Produce json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /v1/health/ping-db [get]
func (h *HealthHandler) PingDB(w http.ResponseWriter, r *http.Request) {
	err := h.DB.PG.PingContext(r.Context())
	if err != nil {
		logger.WarningWithMessage(err, "Failed PING DB!")
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusOK, "DB Postgres is Healthy!")
}
