build:
	go build -o bin/gitbot ./cmd/gitbot/main.go

run: @build
	./bin/gitbot

test:
	go test -v ./...

migrate:
	go run ./cmd/migration/main.go --migrations-path=./migrations
