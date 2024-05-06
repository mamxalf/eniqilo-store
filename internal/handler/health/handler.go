package health

import (
	"github.com/go-chi/chi"
	"github.com/mamxalf/eniqilo-store/http/middleware"
)

type HealthHandler struct {
	JWTMiddleware *middleware.JWT
}

func ProvideHealthHandler(jwt *middleware.JWT) HealthHandler {
	return HealthHandler{
		JWTMiddleware: jwt,
	}
}

func (h *HealthHandler) Router(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/ping", h.Ping)
		})
	})
}
