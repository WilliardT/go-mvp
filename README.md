## Запуск локально

```bash
make env-up
make migrate-up
make env-port-forward
make app-run
```

## Запуск в Docker

```bash
make migrate-up
make go-app-deploy
make ps
```

`go-app-deploy` не накатывает миграции автоматически, их нужно запускать отдельно.

## Swagger

Сгенерировать документацию:

```bash
make swagger-gen
```

Swagger UI:

- `http://localhost:5050/swagger/index.html`

## Конфиг

Основные переменные в `.env` лежат в `.env-example`:

- `HTTP_ADDR`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`
- `POSTGRES_TIMEOUT`
- `LOGGER_LEVEL`
- `TIME_ZONE`
