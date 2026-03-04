# Kubed Makefile — run before pushing to main to ensure quality.
# Similar to Kubernetes: make verify = tests + vet; gate all changes.
SHELL := /bin/bash
BINARY := build/kubed
GOFLAGS ?=
export GOFLAGS

.PHONY: all build test verify vet clean clean-all help

all: verify

# Build the kubed binary (output in build/ to avoid clashing with kubed/ package dir).
build:
	@mkdir -p build
	go build -o $(BINARY) ./cmd/kubed

# Run all tests (unit + agent-query tests).
test:
	go test -v ./...

# Run tests with race detector (CI or pre-push).
test-race:
	go test -race -v ./...

# Verify: vet + test. Run this before pushing to main.
verify: vet test

# Run go vet.
vet:
	go vet ./...

# Clean build artifacts (binary only).
clean:
	rm -f $(BINARY)
	go clean -testcache

# Remove entire build/ directory.
clean-all: clean
	rm -rf build

help:
	@echo "Targets:"
	@echo "  make build     - build binary at $(BINARY)"
	@echo "  make test      - run tests"
	@echo "  make test-race - run tests with race detector"
	@echo "  make vet       - run go vet"
	@echo "  make verify    - vet + test (run before push)"
	@echo "  make clean     - remove binary and test cache"
	@echo "  make clean-all - remove build/ directory"
