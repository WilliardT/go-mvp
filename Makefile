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