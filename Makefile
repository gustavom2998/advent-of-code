fmt:
	go fmt ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

go test:
	go test ./... -race
