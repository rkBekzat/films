postgres:
	docker run --name=films-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

swagger-up:
	swag init --parseDependency --parseInternal --parseDepth 2 -g cmd/main.go 