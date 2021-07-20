#!make
include local/.env
export $(shell sed 's/=.*//' local/.env)

# TODO: fix env (dont work on server)

docker:
	docker build -t bigspawn:music-news .

lint:
	golangci-lint run ./...

test:
	go test -v ./...

upgrade_prod:
	git pull
	docker-compose -f docker-compose-prod.yml rm -s -f news notifier
	docker-compose -f docker-compose-prod.yml up --build -d news notifier

run_parser:
	go run cmd/main.go

run_notifier:
	NOTIFY=true go run cmd/main.go
