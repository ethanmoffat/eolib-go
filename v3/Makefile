GOFMT_FILES := $$(find . -type f -name '*.go' ! -name '*_generated.go' ! -name 'packetmap_generated.go')

build: fmt lint generate
	@go build ./...

test:
	@go test ./...

fmt:
	@files="$(GOFMT_FILES)"; \
	if [ -n "$$files" ]; then gofmt -w $$files; fi

lint:
	@lint_bin="$$(command -v golangci-lint 2>/dev/null || echo "$$(go env GOPATH)/bin/golangci-lint")"; \
	if [ ! -x "$$lint_bin" ]; then \
		echo "golangci-lint not found; install v2.11.4 first (see README.md)"; \
		exit 1; \
	fi; \
	GOCACHE=$${GOCACHE:-/tmp/go-build-v3} GOMODCACHE=$${GOMODCACHE:-/tmp/go-mod-v3} GOLANGCI_LINT_CACHE=$${GOLANGCI_LINT_CACHE:-/tmp/golangci-lint-v3} "$$lint_bin" run ./...

test-cover:
	@go test -coverprofile=c.out ./...
	@-rm c.out

install:
	@go install ./cmd/...

generate:
	@go install ./cmd/protocol-gen-v3
	@go generate .

clean:
	@rm $$(go env GOPATH)/bin/protocol-gen-v3
	@go clean ./...

help:
	@echo "targets:"
	@echo "  build                build the code"
	@echo "  test                 run unit tests"
	@echo "  fmt                  format handwritten Go files"
	@echo "  lint                 run golangci-lint"
	@echo "  test-cover           run unit tests with test coverage"
	@echo "  install              compile and install projects in the cmd directory"
	@echo "  generate             install the code generator under \$$GOPATH/bin and generate code using the default relative paths"
	@echo "  clean                remove the installed protocol generator and clean any build files"
