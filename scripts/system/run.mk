install:  ## Установка зависимостей проекта
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest  
	go install github.com/swaggo/swag/cmd/swag@latest

run:
	bash -c 'set -a; . ./build/local/.env; set +a; go run cmd/codegen-cli/main.go auth-id $(shell pwd)'
