
run:
	go run cmd/main.go

up:
	docker compose up -d

down:
	docker compose down

clear:
	docker compose down -v --rmi all
