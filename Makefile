LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

GOLANGCI_LINT_VERSION ?= v1.61.0

GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint

golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

fmt:
	go fmt ./...

linters: golangci-lint 
	$(GOLANGCI_LINT) run -v --timeout 5m ./...

test:
	go test ./... -coverprofile cover.out

build: fmt linters test build_server build_httpclient build_tcpclient

build_httpclient:
	go build -o bin/httpclient ./cmd/httpclient/main.go

build_tcpclient:
	go build -o bin/tcpclient ./cmd/tcpclient/main.go

build_server:
	go build -o bin/server ./cmd/server/main.go
