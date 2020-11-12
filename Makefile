.DEFAULT_GOAL := help
PROJECTNAME := schemashape

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)
GO = go
GOFMT = gofmt
TIMEOUT = 15
MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
# Project version. If no exist -> changeset id
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
# Go packages
PKGS = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
# Test packages
TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))
# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
M = $(shell printf "\033[34;1m▶\033[0m")
ROOT_CLI_MAIN_FOLDER = cmd/cli

# Tools
$(GOBIN):
	@mkdir -p $@
$(GOBIN)/%: | $(GOBIN) ; $(info $(M) building $(PACKAGE)…)
	tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(GOBIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret
GOLINT = $(GOBIN)/golint # Get linter
$(GOBIN)/golint: PACKAGE=golang.org/x/lint/golint

export GO111MODULE=on

# Commands
## build: Build program binary
.PHONY: build
build: vendor fmt lint ; $(info $(M) building executable...) @
	$(GO) build \
		-tags release \
		-ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' \
		-o $(PROJECTNAME) $(ROOT_CLI_MAIN_FOLDER)/main.go

## vendor: Download the third party libraries
.PHONY: vendor
vendor: ; $(info $(M) downloading third party libraries...) @
	$(GO) mod vendor

## test: Execute the application test
.PHONY: test
test: fmt lint; $(info $(M) running tests...) @
	$(GO) test $(TESTPKGS)

## test-bench: Run benchmarks
.PHONY: test-bench
test-bench: fmt lint; $(info $(M) running bench tests...) @
	$(GO) test -run=__absolutelynothing__ -bench=. $(TESTPKGS)

## test-short: Run only short tests
.PHONY: test-short
test-short: fmt lint; $(info $(M) running short tests...) @
	$(GO) test -short $(TESTPKGS)

## test-verbose: Run tests in verbose mode with coverage reporting
.PHONY: test-verbose
test-verbose: fmt lint; $(info $(M) running verbose tests...) @
	$(GO) test -v $(TESTPKGS)

## test-race: Run tests with race detector
.PHONY: test-race
test-race: fmt lint; $(info $(M) running race tests...) @
	$(GO) test -race $(TESTPKGS)

## lint: Run golint
.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint...) @
	$(GOLINT) -set_exit_status $(PKGS)

## fmt: Run gofmt on all source files
.PHONY: fmt
fmt: | $(GOLINT) ; $(info $(M) running go fmt...) @
	$(GOFMT) -l -w $(SRC)

## simplify: Simplify the go code.
.PHONY: simplify
simplify: | $(GOLINT) ; $(info $(M) simplifying...) @
	$(GOFMT) -s -l -w $(SRC)

## check: Check the code
.PHONY: check
check: ; $(info $(M) checking...) @
	@test -z $(shell $(GOFMT) -l $(ROOT_CLI_MAIN_FOLDER)/main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do bin/golint $${d}; done
	@go vet ${PKGS}

## clean: Clean build files. Runs `go clean` internally.
.PHONY: clean
clean: ; $(info $(M) cleaning...)	@
	@rm -rf $(GOBIN)
	@rm -rf go.sum
	@rm -rf vendor
	@rm $(PROJECTNAME)
	@rm -rf test/tests.* test/coverage.*
	$(GO) clean

## version: Show the project version
.PHONY: version
version:
	@echo $(VERSION)

## help: This message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
