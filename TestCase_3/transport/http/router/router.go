package router

import (
	credit_card_handler "eratani_assesment_test/TestCase_3/internal/handler/credit_card"

	"github.com/go-chi/chi/v5"
)

// DomainHandlers is a struct that contains all domain-specific handlers.
type DomainHandlers struct {
	CreditCardHandler credit_card_handler.CreditCardHandler
}

// Router is the router struct containing handlers.
type Router struct {
	DomainHandlers DomainHandlers
}

// ProvideRouter is the provider function for this router.
func ProvideRouter(domainHandlers DomainHandlers) Router {
	return Router{
		DomainHandlers: domainHandlers,
	}
}

// SetupRoutes sets up all routing for this server.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(rc chi.Router) {
		r.DomainHandlers.CreditCardHandler.Router(rc)
	})
}
