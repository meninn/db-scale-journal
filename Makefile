.PHONY: up down reset logs build run test mysql slowlog

up:            ## sobe mysql + api
	docker compose up -d --build

down:          ## derruba os containers (mantém dados)
	docker compose down

reset:         ## derruba e apaga o volume do mysql (reaplica migrations)
	docker compose down -v

logs:
	docker compose logs -f

build:
	go build -o bin/api ./cmd/api

run:           ## roda a api localmente contra o mysql do compose
	go run ./cmd/api

test:
	go test ./...

mysql:         ## abre um shell mysql
	docker compose exec mysql mysql -uapp -papp journal

slowlog:
	docker compose exec mysql tail -f /var/lib/mysql/slow.log
