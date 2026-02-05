package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "eratani_assesment_test/TestCase_3/internal/domain/credit_card/model"
	"eratani_assesment_test/TestCase_3/transport/http/response"

	"github.com/go-chi/chi/v5"
)

func (h *CreditCardHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, users)
}

func (h *CreditCardHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.WithMessage(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		response.WithError(w, err)
		return
	}
	if user == nil {
		response.WithMessage(w, http.StatusNotFound, "user not found")
		return
	}
	response.WithJSON(w, http.StatusOK, user)
}

func (h *CreditCardHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WithMessage(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.UserService.CreateUser(req)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusCreated, resp)
}

func (h *CreditCardHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.UserService.GetTransactions()
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, transactions)
}

func (h *CreditCardHandler) GetCreditCardTypes(w http.ResponseWriter, r *http.Request) {
	stats, err := h.UserService.GetCreditCardTypes()
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, stats)
}
