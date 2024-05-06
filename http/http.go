package http

import (
	"fmt"
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/docs"
	"github.com/mamxalf/eniqilo-store/http/router"
	"github.com/mamxalf/eniqilo-store/infra"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HTTP is the HTTP server.
type HTTP struct {
	Config *config.Config
	DB     *infra.PostgresConn
	Router router.Router
	mux    *chi.Mux
}

// ProvideHTTP is the provider for HTTP.
func ProvideHTTP(db *infra.PostgresConn, config *config.Config, router router.Router) *HTTP {
	return &HTTP{
		DB:     db,
		Config: config,
		Router: router,
	}
}

// SetupAndServe sets up the server and gets it up and running.
func (h *HTTP) SetupAndServe() {
	h.mux = chi.NewRouter()
	h.setupMiddleware()
	h.setupSwaggerDocs()
	h.setupRoutes()

	h.logServerInfo()

	log.Info().Str("port", h.Config.AppPort).Msg("Starting up HTTP server.")

	server := &http.Server{
		Addr:    ":" + h.Config.AppPort,
		Handler: h.mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Err(err)
	}
}

func (h *HTTP) setupSwaggerDocs() {
	if h.Config.AppEnv == "development" {
		docs.SwaggerInfo.Title = h.Config.AppName
		swaggerURL := fmt.Sprintf("%s/swagger/doc.json", h.Config.AppUrl)
		h.mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))
		log.Info().Str("url", swaggerURL).Msg("Swagger documentation enabled.")
	}
}

func (h *HTTP) setupRoutes() {
	h.Router.SetupRoutes(h.mux)
}

func (h *HTTP) setupMiddleware() {
	h.mux.Use(middleware.Logger)
	h.mux.Use(middleware.Recoverer)
	h.setupCORS()
}

func (h *HTTP) logServerInfo() {
	h.logCORSConfigInfo()
}

func (h *HTTP) logCORSConfigInfo() {
	config := h.Config
	corsHeaderInfo := "CORS Header"
	if config.CorsEnable {
		log.Info().Msg("CORS Headers and Handlers are enabled.")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Credentials: %t", config.CorsAllowCredentials)).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Headers: %s", strings.Join(config.CorsAllowedHeaders, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Methods: %s", strings.Join(config.CorsAllowedMethods, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Origin: %s", strings.Join(config.CorsAllowedOrigins, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Max-Age: %d", config.CorsMaxAgeSeconds)).Msg("")
	} else {
		log.Info().Msg("CORS Headers are disabled.")
	}
}

func (h *HTTP) setupCORS() {
	config := h.Config
	if config.CorsEnable {
		h.mux.Use(cors.Handler(cors.Options{
			AllowCredentials: config.CorsAllowCredentials,
			AllowedHeaders:   config.CorsAllowedHeaders,
			AllowedMethods:   config.CorsAllowedMethods,
			AllowedOrigins:   config.CorsAllowedOrigins,
			MaxAge:           config.CorsMaxAgeSeconds,
		}))
	}
}
