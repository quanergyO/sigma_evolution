run:
	go run cmd/main.go

build:
	mkdir build || true
	go build -o ./build/server.out

compose-run:
	docker-compose up

compose-rebuild:
	docker-compose up --build

dbrun:
	docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5432:5436 -d --rm postgres