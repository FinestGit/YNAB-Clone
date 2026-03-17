# Variables
BINARY_NAME := ynab-clone
MAIN_PACKAGE := ./cmd/api
GO := go
BUILD_DIR := bin

# Go build flags
LDFLAGS := -s -w
BUILD_FLAGS := -ldflags "$(LDFLAGS)"

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the API server
.PHONY: run
run:
	$(GO) run $(MAIN_PACKAGE)

## build: build the API binary
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

## clean: remove build artifacts
.PHONY: tidy
clean:
	rm -rf $(BUILD_DIR)
	$(GO) clean

## tidy: tidy and verify module dependencies
.PHONY: tidy
tidy:
	$(GO) mod tidy
	$(GO) mod verify

## fmt: format all go files
.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: vet
vet:
	$(GO) vet ./...

## test: run all tests
.PHONY: test
test:
	$(GO) test -race ./...

## test/cover: run tests with coverage report
.PHONY: test/cover
test/cover:
	$(GO) test -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

## audit: run quality checks (vet, test, tidy chec)
.PHONY: audit
audit: vet
	$(GO) mod tidy -diff
	$(GO) test -race -vet=off ./...
