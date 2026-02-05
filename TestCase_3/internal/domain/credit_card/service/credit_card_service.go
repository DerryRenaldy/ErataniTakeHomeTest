package service

import (
	model "eratani_assesment_test/TestCase_3/internal/domain/credit_card/model"
	"eratani_assesment_test/TestCase_3/internal/domain/credit_card/repository"
)

type CreditCardService interface {
	GetUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(req model.CreateUserRequest) (model.CreateUserResponse, error)
	GetTransactions() ([]model.TransactionWithUser, error)
	GetCreditCardTypes() ([]map[string]interface{}, error)
}

type CreditCardServiceImpl struct {
	repo repository.CreditCardRepository
}

func ProvideUserService(repo repository.CreditCardRepository) *CreditCardServiceImpl {
	return &CreditCardServiceImpl{repo: repo}
}

func (s *CreditCardServiceImpl) GetUsers() ([]model.User, error) {
	return s.repo.GetAllUser()
}

func (s *CreditCardServiceImpl) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *CreditCardServiceImpl) CreateUser(req model.CreateUserRequest) (model.CreateUserResponse, error) {
	user := &model.User{
		Country:          req.Country,
		CreditCardType:   req.CreditCardType,
		CreditCardNumber: req.CreditCardNumber,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
	}

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return model.CreateUserResponse{}, err
	}

	return model.CreateUserResponse{
		ID:   id,
		Name: req.FirstName + " " + req.LastName,
	}, nil
}

func (s *CreditCardServiceImpl) GetTransactions() ([]model.TransactionWithUser, error) {
	return s.repo.GetTransactions()
}

func (s *CreditCardServiceImpl) GetCreditCardTypes() ([]map[string]interface{}, error) {
	return s.repo.GetCreditCardTypes()
}
