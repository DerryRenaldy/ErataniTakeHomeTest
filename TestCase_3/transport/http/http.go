package http

import (
	"eratani_assesment_test/TestCase_3/configs"
	"eratani_assesment_test/TestCase_3/database"
	"eratani_assesment_test/TestCase_3/transport"
	"eratani_assesment_test/TestCase_3/transport/http/response"
	"eratani_assesment_test/TestCase_3/transport/http/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type HTTP struct {
	Config *configs.Config
	DB     *database.PostgresConn
	State  transport.ServerState
	Router router.Router
	Mux    *chi.Mux
}

func ProvideHTTP(config *configs.Config, db *database.PostgresConn, router router.Router) *HTTP {
	return &HTTP{
		Config: config,
		DB:     db,
		Router: router,
	}
}

func (h *HTTP) GetState() transport.ServerState {
	return h.State
}

func (h *HTTP) SetState(state transport.ServerState) {
	h.State = state
}

func (h *HTTP) GetGracePeriodSeconds() int64 {
	return h.Config.Server.Shutdown.GracePeriodSeconds
}

func (h *HTTP) GetCleanupPeriodSeconds() int64 {
	return h.Config.Server.Shutdown.CleanupPeriodSeconds
}

func (h *HTTP) SetupAndServe() {
	h.Mux = chi.NewRouter()
	h.setupMiddleware()
	h.setupRoutes()
	log.Info().Msgf("Starting server on port %s", h.Config.Server.Port)
	if err := http.ListenAndServe(":"+h.Config.Server.Port, h.Mux); err != nil {
		log.Info().Msgf("Server stopped: %s", err)
	}
}

func (h *HTTP) setupMiddleware() {
	h.Mux.Use(middleware.Logger)
	h.Mux.Use(middleware.Recoverer)
	h.Mux.Use(h.serverStateMiddleware)
}

func (h *HTTP) serverStateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch h.State {
		case transport.ServerStateReady:
			log.Info().Msg("SERVER IS READY")
			next.ServeHTTP(w, r)
		case transport.ServerStateInGracePeriod:
			log.Warn().Msg("SERVER IS IN GRACE PERIOD")
			next.ServeHTTP(w, r)
		case transport.ServerStateInCleanupPeriod:
			log.Warn().Msg("SERVER IS IN CLEANUP PERIOD")
			response.WithPreparingShutdown(w)
		}
	})
}

func (h *HTTP) setupRoutes() {
	h.Mux.Get("/health", h.HealthCheck)
	h.Router.SetupRoutes(h.Mux)

}

func (h *HTTP) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.WithMessage(w, http.StatusOK, "OK")
}
