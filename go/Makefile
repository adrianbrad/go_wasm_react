
GOCMD=go
GOBUILD=$(GOCMD) build

# Binary names
BINARY_NAME=main.wasm

build:
	GOOS=js GOARCH=wasm $(GOBUILD) -o $(BINARY_NAME) -v