#!make
lint:
	golangci-lint run ./...

test:
	go test -v ./...

run_parser:
	go run cmd/main.go

run_notifier:
	NOTIFY=true go run cmd/main.go
