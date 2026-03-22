include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d go-mvp-postgres

env-down:
	@docker compose down go-mvp-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. (y/N): " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down go-mvp-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Имя миграции не может быть пустым. Пример: make migrate-create name=create_users_table"; \
		exit 1; \
	fi ;\
	docker compose run --rm go-mvp-postgres-migrate \
		create \
	  	-ext sql \
		-dir /migrations \
		-seq "$(name)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Действие не может быть пустым. Пример: make migrate-action action=up"; \
		exit 1; \
	fi ;\
	docker compose run --rm go-mvp-postgres-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@go-mvp-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		$(action)

app-run:
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	go mod tidy && \
	go run cmd/go-mvp-app/main.go