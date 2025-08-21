include .env
GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOIMPORTS := $(GOPATH_BIN)/goimports
GO_PACKAGES = $(shell go list ./... | grep -v vendor)
PACKAGE_BASE := github/aryan-go/food_ordering_go

UP_MIGRATION_FILE = database/migrations/000004_create_food_menu_table.up.sql
DOWN_MIGRATION_FILE = database/migrations/000001_create_food_go_database.down.sql

.PHONY: help vendor build run dev lint format clean

help:
	@echo "Aryan make help"
	@echo ""
	@echo "vendor: Downloads the dependencies in the vendor folder"
	@echo "build: Builds the binary of the server"
	@echo "run: Runs the binary of the server"
	@echo "format: Formats the code using fmt and golangci-lint"
	@echo "clean: Removes the vendor directory and binary"


run:
	@${GO} run command/main.go

install-goimports:
	@echo "=====> Installing formatter..."
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

format: install-goimports
	@echo "=====> Formatting code..."
	@$(GOIMPORTS) -l -w -local ${PACKAGE_BASE} $(SRC)
	@echo "Format successful"

## verify: Run format and lint checks
verify: verify-format

## verify-format: Verify the format
verify-format: install-goimports
	@echo "=====> Verifying format..."
	$(if $(shell $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC}), @echo ERROR: Format verification failed! && $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC} && exit 1)

apply-migration:
	@echo "Applying migration..."
	@echo "DB_HOST: $(db_host)"
	@echo "DB_PORT: $(db_port)"
	@echo "DB_USER: $(db_user)"
# 	@echo "DB_PASS: $(db_password)"
	@echo "DB_NAME: $(db_name)"
	@echo "MIGRATIONS_DIR: $(MIGRATIONS_DIR)"
	migrate -path $(MIGRATIONS_DIR) -database "mysql://$(db_user):$(db_password)@tcp($(db_host):$(db_port))/$(db_name)" up

rollback-migration:
	@echo "Rolling back migration..."
	migrate -path $(MIGRATIONS_DIR) -database "mysql://$(db_user):$(db_password)@tcp($(db_host):$(db_port))/$(db_name)" down

test:
	go test -v ./package/middlewares
