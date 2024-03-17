postgres:
	docker run --name=films-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
