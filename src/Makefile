.PHONY: up down run test

up: 
	docker compose up --build -d

down:
	docker compose down

run:
	go build -o bin/main main.go

test:
	go test -v ./test