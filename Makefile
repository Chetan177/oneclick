
# Define the source files
# CLI_SRC = $(SRC_DIR)/cli/main.go
SERVER_SRC = $(SRC_DIR)/server/main.go

# Define the output binaries
# CLI_BIN = $(BIN_DIR)/cli
SERVER_BIN = $(BIN_DIR)/server

# Define the build targets
build: clean $(CLI_BIN) $(SERVER_BIN)

# $(CLI_BIN): $(CLI_SRC)
# 	@mkdir -p $(BIN_DIR)
# 	CGO_ENABLED=0 go build -o $@ $<

$(SERVER_BIN): $(SERVER_SRC)
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -o $@ $<

clean:
	rm -rf $(BIN_DIR)

.PHONY: build clean

help:
	@echo "Makefile targets:"
	@echo "  build   - Build all binaries (CLI and Server)"
	@echo "  clean   - Remove all binaries"
	@echo "  help    - Display this help message"