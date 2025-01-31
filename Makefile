# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=stripper
GITHUB_USERNAME=ncecere

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build clean test coverage lint install uninstall

all: lint test build

build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	rm -rf output/
	rm -f *.db
	rm -f coverage.out

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

lint:
	golangci-lint run

install:
	$(GOCMD) install -v ./...

uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

# Docker targets
docker-build:
	docker build -t $(GITHUB_USERNAME)/$(BINARY_NAME) .

docker-push:
	docker push $(GITHUB_USERNAME)/$(BINARY_NAME)

# Development targets
dev: build
	./$(BINARY_NAME)

run: build
	./$(BINARY_NAME)

# Help target
help:
	@echo "Available targets:"
	@echo "  make          - Run lint, test, and build"
	@echo "  make build    - Build the binary"
	@echo "  make clean    - Remove binary and artifacts"
	@echo "  make test     - Run tests"
	@echo "  make coverage - Generate test coverage report"
	@echo "  make lint     - Run linter"
	@echo "  make install  - Install binary to GOPATH"
	@echo "  make dev      - Build and run for development"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-push  - Push Docker image to registry"
