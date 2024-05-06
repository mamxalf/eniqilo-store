package health

import (
	"github.com/mamxalf/eniqilo-store/shared/response"
	"net/http"
)

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
