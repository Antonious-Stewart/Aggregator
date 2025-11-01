change_dir:
	cd cmd/aggregator

migrate_up:
	goose up

migrate_down:
	goose down

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

build: change_dir
	 go build -o bin/main main.go

run: fmt vet test change_dir
	go run main.go

air: fmt vet test
	air
