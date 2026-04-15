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
	@echo "  test-cover-all       test coverage for all versions of the code"
	@echo "  test-cover           run unit tests with test coverage"
	@echo "  install              compile and install projects in the cmd directory"
	@echo "  generate             install the code generator under \$$GOPATH/bin and generate code using the default relative paths"
	@echo "  clean                remove the installed protocol generator and clean any build files"
	@echo ""
	@echo "append '-v3' to a target to do it for version 3"
