TARGET=oju
BUILD_DEBUG_DIR=build/debug

run:
	@./$(BUILD_DEBUG_DIR)/$(TARGET)

build:
	@go build -o $(BUILD_DEBUG_DIR)/$(TARGET) ./cmd/oju

.PHONY: build
