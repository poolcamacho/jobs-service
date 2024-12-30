swagger:
	swag init --dir ./cmd,./internal/transport,./internal/domain --output ./docs

run:
	go run cmd/main.go

test:
	go test ./... -coverprofile=coverage.out