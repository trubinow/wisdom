.PHONY: build-client
build:
	 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "build/client" -ldflags '-w' cmd/client.go
.PHONY: build-server
build:
	 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "build/server" -ldflags '-w' cmd/server.go
.PHONY: lint
lint:
	golangci-lint run
