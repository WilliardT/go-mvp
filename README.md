## Локальный запуск

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

Важно: `go-app-deploy` не накатывает миграции автоматически, их нужно запускать отдельно.

## Полезные команды Makefile

## Конфиг

Основные нужные переменные лежат в `.env-example`:
- `HTTP_ADDR`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`
- `POSTGRES_TIMEOUT`
- `LOGGER_LEVEL`
- `TIME_ZONE`
