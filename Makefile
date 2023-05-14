CURRENT_DIR=$(shell pwd)

PSQL_CONTAINER_NAME=postgres-container
PROJECT_NAME=back-end-app
PSQL_URI ?=postgres://postgres:postgres@localhost:5432/${PROJECT_NAME}?sslmode=disable

IMAGE_TAG ?= latest
REGISTRY_TAG ?= registry.gitlab.com/greatsoft1/xif-new/xif-backend

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: swag_init
swag_init:
	swag init -g internal/transport/handlers/server.go -o api/openapi  --outputTypes "go,json" --overridesFile .swaggo

.PHONY: mockgen
mockgen:
	mockgen -package mockdb -destination internal/core/repository/psql/mock/store.go github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc Querier

.PHONY: image_build
image_build:
	docker build -t ${PROJECT_NAME}:${IMAGE_TAG} .

.PHONY: tag_image
tag_image:
	docker tag ${PROJECT_NAME}:${IMAGE_TAG} ${REGISTRY_TAG}:${IMAGE_TAG}

.PHONY: push_image
push_image:
	docker push ${REGISTRY_TAG}:${IMAGE_TAG}

.PHONY: run_dokcer_image
run_docker_image:
	docker run --network=host -e APPLICATION_MODE=dev -e PSQL_URI=${PSQL_URI} -p 8080:8080 ${PROJECT_NAME}:latest

.PHONY: dev_environment_remove
dev_environment_remove:
	docker compose -f docker-compose.dev.yml down --volumes

.PHONY: createdb
createdb:
	docker exec -it ${PSQL_CONTAINER_NAME} createdb -U postgres ${PROJECT_NAME}

execdb:
	docker exec -it ${PSQL_CONTAINER_NAME} psql -U postgres ${PROJECT_NAME}

.PHONY: dropdb
dropdb:
	docker exec -it ${PSQL_CONTAINER_NAME} dropdb -U postgres ${PROJECT_NAME}

cleandb:
	docker exec -it ${PSQL_CONTAINER_NAME} psql -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" ${PSQL_URI}

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database ${PSQL_URI} -verbose up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database ${PSQL_URI} -verbose down 1

.PHONY: migrate_force
migrate_force:
	migrate -path migrations -database ${PSQL_URI} force ${MIGRATE_FORCE_VERSION}

.PHONY: migrate_file
migrate_file:
	migrate create -ext sql -dir ./migrations -seq ${FILE_NAME}

.PHONY: run
run:
	docker compose -f docker-compose.dev.yml up -d --build

.PHONY: stop
stop:
	docker compose -f docker-compose.dev.yml down

.PHONY: restart
restart:
	docker compose -f docker-compose.dev.yml down
	docker compose -f docker-compose.dev.yml up -d --build

.PHONY: remove
remove:
	docker compose -f docker-compose.dev.yml down --volumes

.PHONY: logs
logs:
	docker compose -f docker-compose.dev.yml logs -f $(name)

seed:
	go run ./mocks/db/main.go
