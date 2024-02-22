docker-up:
	docker compose up --build -d

docker-down:
	docker compose down

go-run:
	go run ./src/main.go
	