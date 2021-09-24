format:
	gofmt -s -w .

build:
	go build ./...

test:
	go clean -testcache
	go test ./... -v

.PHONY: openapihttp
openapihttp:
	oapi-codegen -generate types -o internal/app/ports/openapi/todotypes.go -package openapi api/openapi/todo.yml
	oapi-codegen -generate chi-server -o internal/app/ports/openapi/todoapi.go -package openapi api/openapi/todo.yml
	oapi-codegen -generate types -o internal/client/ports/openapi/todotypes.go -package openapi api/openapi/todo.yml
	oapi-codegen -generate client -o internal/client/ports/openapi/todoapi.go -package openapi api/openapi/todo.yml

