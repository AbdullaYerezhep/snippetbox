PROJECT_NAME := bank-account

.PHONY: help
help: ## Show this help.
	@echo "Choose a command run in "$(PROJECT_NAME)":"
	@echo ""
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo "  db/start        Start database"
	@echo "  db/migrate      Run database migrations"
	@echo "  db/rollback     Rollback database migrations"
	@echo "  run             Run application"
	@echo "  test            Run tests"
	@echo ""

.PHONY: db/start
db/start:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DATABASE=snippetbox -d postgres:latest

.PHONY: db/migrate
db/migrate:
	migrate -path db/migrations -database "postgresql://admin:admin@localhost:5432/snippetbox?sslmode=disable" -verbose up

.PHONY: db/rollback
db/rollback:
	migrate -path db/migrations -database "postgresql://admin:admin@localhost:5432/snippetbox?sslmode=disable" -verbose down

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v ./...