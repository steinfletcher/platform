test:
	./scripts/test.sh

postgres-dev:
	docker stop platform || true && docker rm platform || true
	docker run -d -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=platform -p 15432:5432 --name platform postgres:9.6
	sleep 4
	cd migrations && goose postgres "host=localhost port=15432 user=postgres password=postgres dbname=platform sslmode=disable" up
