include .env
export

service-run :
	@go run main.go

deploy:
	docker compose up -d application

un-deploy:
	docker compose down application

migrate-up :
	@migrate -path migration-docker -database ${CONN_STRING} up


migrate-down :
	@migrate -path migration-docker -database ${CONN_STRING} down

migrate-version :
	@migrate -path migration-docker -database ${CONN_STRING} version

postgres-up:
	docker-compose -f docker-compose.practice.yaml up -d postgres

run-http-app:
	docker-compose -f docker-compose.practice.yaml up -d go-app