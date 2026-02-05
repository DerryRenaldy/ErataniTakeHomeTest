dev: generate
	go run github.com/air-verse/air

build:
	go build -o bin/server .

run: build
	./bin/server

test:
	go test -v ./...

wire:
	wire ./...

deps:
	go mod tidy

generate:
	go generate ./...

# =============================================================================
# CLIENT COMMANDS - Run with: make <target>
# =============================================================================

# Base URL for the API
BASE_URL ?= http://localhost:8080
# Base URL for the docker tests
# BASE_URL ?= http://localhost:8090

# JQ is a command-line JSON processor.
JQ ?= jq

# Health check
.PHONY: health
health:
	@echo "=== Health Check ==="
	@curl -s $(BASE_URL)/health
	@echo ""

# Users - Get all
.PHONY: get-users
get-users:
	@echo "=== Get All Users ==="
	@curl -s $(BASE_URL)/v1/users | $(JQ)

# Users - Get by ID - 3c
.PHONY: get-user
get-user:
	@echo "=== Get User by ID ==="
	@if [ -z "$(id)" ]; then \
		echo "Usage: make get-user id=1"; \
	else \
		curl -s $(BASE_URL)/v1/users/$(id) | $(JQ); \
	fi

# Users - Create - 3d
.PHONY: create-user
create-user:
	@echo "=== Create User ==="
	@curl -s -X POST $(BASE_URL)/v1/users \
		-H "Content-Type: application/json" \
		-d '{"country":"Indonesia","credit_card_type":"Visa","credit_card_number":"4111111111111111","first_name":"John","last_name":"Doe"}' | $(JQ)

# Transactions - Get all
.PHONY: get-transactions
get-transactions:
	@echo "=== Get All Transactions ==="
	@curl -s $(BASE_URL)/v1/transactions | $(JQ)

# Transactions - Get credit card types
.PHONY: get-credit-card-types
get-credit-card-types:
	@echo "=== Get Credit Card Types ==="
	@curl -s $(BASE_URL)/v1/transactions/credit-card-types | $(JQ)

.PHONY: dev build run test wire deps generate health get-users get-user create-user get-transactions get-credit-card-types client-test