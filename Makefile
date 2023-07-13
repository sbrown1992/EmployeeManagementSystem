all:
	make verify

verify:
	go fmt ./...
	go mod tidy
	go vet ./...
	go test ./...
