GOFMT_FILES := $$(find . -type f -name '*.go' ! -name '*_generated.go' ! -name 'packetmap_generated.go' ! -path './v3/*')

build-all: build build-v3

build: generate
	@go build ./...

build-v3:
	@make -C v3 build

test-all: test test-v3

test:
	@go test ./...

test-v3:
	@make -C v3 test

fmt-all: fmt fmt-v3

fmt:
	@files="$(GOFMT_FILES)"; \
	if [ -n "$$files" ]; then gofmt -w $$files; fi

fmt-v3:
	@make -C v3 fmt

lint-all: lint lint-v3

lint:
	@GOCACHE=$${GOCACHE:-/tmp/go-build} GOMODCACHE=$${GOMODCACHE:-/tmp/go-mod} GOLANGCI_LINT_CACHE=$${GOLANGCI_LINT_CACHE:-/tmp/golangci-lint} golangci-lint run ./...

lint-v3:
	@make -C v3 lint

test-cover-all: test-cover test-cover-v3

test-cover:
	@go test -coverprofile=c.out ./...
	@-rm c.out

test-cover-v3:
	@make -C v3 test-cover

install:
	@go install ./cmd/...

install-v3:
	@make -C v3 install

generate:
	@go install ./cmd/protocol-gen
	@go generate .

generate-v3:
	@make -C v3 generate

clean:
	@rm $$(go env GOPATH)/bin/protocol-gen
	@go clean ./...

clean-v3:
	@make -C v3 clean

help:
	@echo "targets:"
	@echo "  build-all            build all versions of the code"
	@echo "  build                build the code"
	@echo "  test-all             test all versions of the code"
	@echo "  test                 run unit tests"
	@echo "  fmt-all              format handwritten Go files for all versions"
	@echo "  fmt                  format handwritten Go files"
	@echo "  lint-all             run golangci-lint for all versions"
	@echo "  lint                 run golangci-lint"
	@echo "  test-cover-all       test coverage for all versions of the code"
	@echo "  test-cover           run unit tests with test coverage"
	@echo "  install              compile and install projects in the cmd directory"
	@echo "  generate             install the code generator under \$$GOPATH/bin and generate code using the default relative paths"
	@echo "  clean                remove the installed protocol generator and clean any build files"
	@echo ""
	@echo "append '-v3' to a target to do it for version 3"
