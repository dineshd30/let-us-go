# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean

# Build target
.PHONY: build
build:
	$(GOBUILD) -o ./bin/todos ./cmd/todos

# Run target
.PHONY: run
run: clean build
	./bin/todos

# Test target
.PHONY: test
test:
	$(GOTEST) -v ./...

# Clean target
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf ./bin

# Build proto
.PHONY: proto
proto:
	protoc \
	--go_out=./todos \
	--go_opt=paths=source_relative \
	--go-grpc_out=./todos \
	--go-grpc_opt=paths=source_relative \
	./todos.proto

