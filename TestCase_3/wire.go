//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"eratani_assesment_test/TestCase_3/configs"
	"eratani_assesment_test/TestCase_3/database"
	"eratani_assesment_test/TestCase_3/internal/domain/credit_card/repository"
	"eratani_assesment_test/TestCase_3/internal/domain/credit_card/service"
	credit_card_handler "eratani_assesment_test/TestCase_3/internal/handler/credit_card"
	"eratani_assesment_test/TestCase_3/transport/http"
	"eratani_assesment_test/TestCase_3/transport/http/router"
)

// Config provider
var configGen = wire.NewSet(
	configs.Get,
)

// Database provider
var databaseGen = wire.NewSet(
	database.ProvidePostgresConn,
)

// Repository provider + binding
var repoGen = wire.NewSet(
	repository.ProvideUserRepository,
	wire.Bind(new(repository.CreditCardRepository), new(*repository.CreditCardRepositoryPostgres)),
)

// Service provider + binding
var svcGen = wire.NewSet(
	service.ProvideUserService,
	wire.Bind(new(service.CreditCardService), new(*service.CreditCardServiceImpl)),
)

// Handler provider
var handlerGen = wire.NewSet(
	credit_card_handler.ProvideUserHandler,
)

// Domain handlers provider
var domainHandlersGen = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
)

// Router provider
var routerGen = wire.NewSet(
	router.ProvideRouter,
)

// HTTP provider
var httpGen = wire.NewSet(
	http.ProvideHTTP,
)

// InitializeHTTPService initializes the HTTP service with all dependencies
func InitializeHTTPService() *http.HTTP {
	wire.Build(
		configGen,
		databaseGen,
		repoGen,
		svcGen,
		handlerGen,
		domainHandlersGen,
		routerGen,
		httpGen,
	)
	return &http.HTTP{}
}
