BIN_ROTATION = "./bin"

build:
	go build -v -o ${BIN_ROTATION} ./cmd
lint:
	golangci-lint run ./...
test:
	go test -v -count=1 -race -gcflags=-l -timeout=30s ./...
