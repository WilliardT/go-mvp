include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d go-mvp-postgres

env-down:
	docker compose down go-mvp-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. (y/N): " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down go-mvp-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Имя миграции не может быть пустым. Пример: make migrate-create name=create_users_table"; \
		exit 1; \
	fi \

	docker compose run --rm go-mvp-postgres-migrate \
		create \
	  	-ext sql \
		-dir /migrations \
		-seq "$(name)"

