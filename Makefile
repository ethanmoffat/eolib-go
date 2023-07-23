build: generate
	@go build ./...

test:
	@go test -v ./...

install:
	@go install ./cmd/...

generate:
	@go install ./cmd/protocol-gen
	@go generate ./...

clean:
	@rm $$(go env GOPATH)/bin/protocol-gen
	@go clean ./...

help:
	@echo "targets:"
	@echo "  build                build the code"
	@echo "  test                 run unit tests"
	@echo "  install              compile and install projects in the cmd directory"
	@echo "  generate             install the code generator under \$$GOPATH/bin and generate code using the default relative paths"
	@echo "  clean                remove the installed protocol generator and clean any build files"