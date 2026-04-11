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
	@migrate -path migrations -database ${CONN_STRING} up


migrate-down :
	@migrate -path migrations -database ${CONN_STRING} down

migrate-version :
	@migrate -path migrations -database ${CONN_STRING} version
