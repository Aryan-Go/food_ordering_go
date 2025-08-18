include .env
GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOIMPORTS := $(GOPATH_BIN)/goimports
GO_PACKAGES = $(shell go list ./... | grep -v vendor)
PACKAGE_BASE := github/aryan-go/food_ordering_go

# DB_HOST = $(shell grep -A6 "^db:" config.yaml | grep "host:" | head -1 | cut -d'"' -f2)
# DB_PORT = $(shell grep -A6 "^db:" config.yaml | grep "port:" | head -1 | awk '{print $$2}')
# DB_USER = $(shell grep -A6 "^db:" config.yaml | grep "user:" | head -1 | cut -d'"' -f2)
# DB_PASS = $(shell grep -A6 "^db:" config.yaml | grep "password:" | head -1 | cut -d'"' -f2)
# DB_NAME = $(shell grep -A6 "^db:" config.yaml | grep "db_name:" | head -1 | cut -d'"' -f2)

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

vendor:
	@${GO} mod tidy
	@${GO} mod vendor
	@echo "Vendor downloaded successfully"

build:
	@${GO} build -o aryan ./command/main.go
	@echo "Binary built successfully"

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
	@echo "DB_HOST: $(db_host)"
	@echo "DB_PORT: $(db_port)"
	@echo "DB_USER: $(db_user)"
	@echo "DB_PASS: $(db_password)"
	mysql -h $(db_host) -P $(db_port) -u $(db_user) -p$(db_password) < $(UP_MIGRATION_FILE)

rollback-migration:
	@echo "Rolling back migration..."
	mysql -h $(db_host) -P $(db_port) -u $(db_user) -p$(db_password) $(db_database) < $(DOWN_MIGRATION_FILE)

test:
	go test -v ./package/middlewares
