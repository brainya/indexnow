.PHONY: generate-go-client
generate-go-client:
	go generate ./...

.PHONY: install-deps
install-deps:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
