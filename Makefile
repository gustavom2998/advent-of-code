fmt:
	go fmt ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

test:
	go test ./... -race

run:
	go run "$(year)/$(day)/main.go"
