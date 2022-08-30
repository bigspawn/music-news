#!make
lint:
	@go vet ./...

test:
	go test -v ./...

run_parser:
	go run cmd/main.go

run_notifier:
	NOTIFY=true go run cmd/main.go

clean:
	@go mod tidy -v
	@go clean -testcache ./...
