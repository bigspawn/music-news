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
