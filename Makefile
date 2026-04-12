include .env
export

run-http-app:
	docker run -p 8080:8080 docker-practice:latest 

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
	docker run -v ./out/pg_data:/var/lib/postgresql -p 5432:5432 -e POSTGRES_PASSWORD="0710" -d postgres:18-bookworm