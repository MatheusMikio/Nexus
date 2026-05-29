.PHONY: default run run-with-docs build test docs clean fmt tidy

APP_NAME=nexus-api
ENTRYPOINT=./cmd/api

default: run

run:
	@go run $(ENTRYPOINT)

run-with-docs: docs run

build:
	@go build -o bin/$(APP_NAME) $(ENTRYPOINT)

test:
	@go test ./...

docs:
	@go run github.com/swaggo/swag/cmd/swag@latest init -g main.go -d cmd/api,internal/handler,internal/domain/dtos/auth,internal/domain/dtos/user,internal/domain/dtos/goal,internal/domain/dtos/task,internal/domain/models,internal/response -o docs --parseInternal
	@echo "Scalar docs: http://localhost:8080/api/v1/docs"
	@echo "OpenAPI JSON: http://localhost:8080/api/v1/openapi.json"

fmt:
	@go fmt ./...

tidy:
	@go mod tidy

clean:
	@rm -rf bin
	@rm -f docs/docs.go docs/swagger.json docs/swagger.yaml
