up:
	docker compose  up --build -d --force-recreate
	docker compose logs -f

down:
	docker compose down

run: mod
	go run ./cmd/app

mod:
	go mod tidy

mod-update:
	go get -u all
	go mod tidy