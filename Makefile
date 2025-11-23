APP_DIR=cmd/aggregator

change_dir:
	@echo "Switching to $(APP_DIR)"

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
	 cd $(APP_DIR) && go build -o ../../bin/main main.go

run: fmt vet test change_dir
	cd $(APP_DIR) && go run main.go

air: fmt vet test
	cd $(APP_DIR) && air