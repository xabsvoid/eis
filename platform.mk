export LOCAL_BIN:=$(CURDIR)/bin

# golang

export GO111MODULE=on
export GOBIN:=$(LOCAL_BIN)

test:
	go test ./...

# lint

GOLANGCI_LINT_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_LINT_VER:=1.63.4
install-golangci-lint:
ifeq ($(wildcard $(GOLANGCI_LINT_BIN)),)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VER)
endif

lint: install-golangci-lint
	$(GOLANGCI_LINT_BIN) run ./...

# mock

MOCKERY_BIN:=$(LOCAL_BIN)/mockery
MOCKERY_VER:=2.53.4
install-mockery:
ifeq ($(wildcard $(MOCKERY_BIN)),)
	go install github.com/vektra/mockery/v2@v$(MOCKERY_VER)
endif

mock: install-mockery
	$(MOCKERY_BIN)

# api

OAPI_CODEGEN_BIN:=$(LOCAL_BIN)/oapi-codegen
OAPI_CODEGEN_VER:=2.4.1
install-oapi-codegen:
ifeq ($(wildcard $(OAPI_CODEGEN_BIN)),)
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v$(OAPI_CODEGEN_VER)
endif

api: install-oapi-codegen
	$(OAPI_CODEGEN_BIN) --config=.oapi-codegen-config.yaml api/v1/openapi.json
