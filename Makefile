include platform.mk

-include .env
export

export DATABASE_URL = postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down

build-eis:
	go build -o $(LOCAL_BIN) ./cmd/eis

build: build-eis

migrate-create:
	docker compose run --rm migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	docker compose run --rm migrate \
	  -path=/migrations \
	  -database "$(DATABASE_URL)" up

migrate-down:
	docker compose run --rm migrate \
	  -path=/migrations \
	  -database "$(DATABASE_URL)" down 1
