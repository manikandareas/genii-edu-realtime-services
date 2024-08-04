CMD_DIR := ./cmd/web
BUILD_DIR := ./tmp
BINARY_NAME := realtime-services

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags "-s" -v -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

dev:
	go run $(CMD_DIR)
