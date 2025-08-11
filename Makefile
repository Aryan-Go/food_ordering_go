# hello:
# 	@echo "Hello"

# test:
# 	 go test -v ./package/middlewares

# run:
# 	go run command/main.go

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOIMPORTS := $(GOPATH_BIN)/goimports
GO_PACKAGES = $(shell go list ./... | grep -v vendor)
PACKAGE_BASE := github/aryan-go/food_ordering_go

DB_HOST = $(shell grep -A6 "^db:" config.yaml | grep "host:" | head -1 | cut -d'"' -f2)
DB_PORT = $(shell grep -A6 "^db:" config.yaml | grep "port:" | head -1 | awk '{print $$2}')
DB_USER = $(shell grep -A6 "^db:" config.yaml | grep "user:" | head -1 | cut -d'"' -f2)
DB_PASS = $(shell grep -A6 "^db:" config.yaml | grep "password:" | head -1 | cut -d'"' -f2)
DB_NAME = $(shell grep -A6 "^db:" config.yaml | grep "db_name:" | head -1 | cut -d'"' -f2)

UP_MIGRATION_FILE = database/migrations/000001_init_schema.up.sql
DOWN_MIGRATION_FILE = database/migrations/000001_init_schema.down.sql

.PHONY: help vendor build run dev lint format clean

help:
	@echo "Aryan make help"
	@echo ""
	@echo "vendor: Downloads the dependencies in the vendor folder"
	@echo "build: Builds the binary of the server"
	@echo "run: Runs the binary of the server"
	@echo "dev: Combines build and run commands"
	@echo "lint: Lints the code using vet and golangci-lint"
	@echo "format: Formats the code using fmt and golangci-lint"
	@echo "clean: Removes the vendor directory and binary"

vendor:
	@${GO} mod tidy
	@${GO} mod vendor
	@echo "Vendor downloaded successfully"

build:
	@${GO} build -o aryan ./command/main.go
	@echo "Binary built successfully"

run:
	@${GO} run command/main.go

dev:
	@$(GOPATH_BIN)/air -c .air.toml

install-golangci-lint:
	@echo "=====> Installing golangci-lint..."
	@curl -sSfL \
	 	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	 	sh -s -- -b $(GOPATH_BIN) v1.62.2

lint: install-golangci-lint
	@$(GO) vet $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run -c golangci.yaml
	@echo "Lint successful"

install-goimports:
	@echo "=====> Installing formatter..."
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

format: install-goimports
	@echo "=====> Formatting code..."
	@$(GOIMPORTS) -l -w -local ${PACKAGE_BASE} $(SRC)
	@echo "Format successful"

## verify: Run format and lint checks
verify: verify-format lint

## verify-format: Verify the format
verify-format: install-goimports
	@echo "=====> Verifying format..."
	$(if $(shell $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC}), @echo ERROR: Format verification failed! && $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC} && exit 1)

clean:
	@rm -f aryan
	@rm -rf vendor/
	@echo "Clean successful"

install-air:
	@echo "Make sure your GOPATH and GOPATH_BIN is set"
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH_BIN)
	@echo "Air installed successfully"	

apply-migration:
	@echo "Applying migration..."
	@echo "DB_HOST: $(DB_HOST)"
	@echo "DB_PORT: $(DB_PORT)"
	@echo "DB_USER: $(DB_USER)"
	@echo "DB_PASS: $(DB_PASS)"
	@echo "DB_NAME: $(DB_NAME)"
	PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f $(UP_MIGRATION_FILE)

rollback-migration:
	@echo "Rolling back migration..."
	PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f $(DOWN_MIGRATION_FILE)