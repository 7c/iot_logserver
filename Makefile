.PHONY: build clean

# Set the output directory
OUT_DIR := bin/

# Set the target output file
TARGET := $(OUT_DIR)iot_logserver

# Default target
all: build

# Build wasm target
build:
	@echo "Building iot_logserver..."
	@go build -o $(TARGET) *.go

# Clean target
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUT_DIR)

# Create bin directory if it doesn't exist
$(shell mkdir -p $(OUT_DIR))
