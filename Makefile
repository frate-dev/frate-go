GO ?= go
GOBIN ?= $(shell $(GO) env GOBIN)
GOPATH ?= $(shell $(GO) env GOPATH)
GOMOD := $(GO) mod
GOTEST := $(GO) test
GOFMT := gofmt
GOLINT := golangci-lint

APP_NAME = frate-go
BIN_DIR = ./bin
MAIN_SRC = ./main.go
BUILD_DIR = ./output

TEST_FLAGS ?= -v -cover

.PHONY: all
all: clean test build lint

.PHONY: run
run: $(MAIN_SRC)
	$(GO) run $(MAIN_SRC)

.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) . 

.PHONY: install
install: 
	@echo "Installing $(APP_NAME) to $(GOBIN)..."
	$(GO) install .

.PHONY: test
test: 
	@echo "Running tests..."
	$(GOTEST) ./... $(TEST_FLAGS)

.PHONY: lint
lint:
	@echo "Running linter..."
	$(GOLINT) run

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w $(GOFILES)

.PHONY: mod-tidy
mod-tidy:
	$(GOMOD) tidy

.PHONY: mod-download
mod-download:
	$(GOMOD) download

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR) $(BIN_DIR)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

