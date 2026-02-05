package handler

import (
	"eratani_assesment_test/TestCase_3/internal/domain/credit_card/service"

	"github.com/go-chi/chi/v5"
)

type CreditCardHandler struct {
	UserService *service.CreditCardServiceImpl
}

func ProvideUserHandler(svc *service.CreditCardServiceImpl) CreditCardHandler {
	return CreditCardHandler{UserService: svc}
}

func (h *CreditCardHandler) Router(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetUsers)
		r.Post("/", h.CreateUser)
		r.Get("/{id}", h.GetUserByID)
	})

	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", h.GetTransactions)
		r.Get("/credit-card-types", h.GetCreditCardTypes)
	})
}
