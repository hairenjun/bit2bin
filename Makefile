# Binary name
BINARY_NAME=bit2bin
# Output directory
BUILD_DIR=build

# Build flags for static linking and smaller binary size
# -s: disable symbol table
# -w: disable DWARF generation
LDFLAGS=-ldflags="-s -w -extldflags=-static"

.PHONY: all clean release linux-amd64 linux-arm64 windows-amd64 windows-arm64

release: linux-amd64 linux-arm64 windows-amd64 windows-arm64
	@echo "Building all platform binaries..."
	@$(MAKE) clean
	@$(MAKE) linux-amd64
	@$(MAKE) linux-arm64
	@$(MAKE) windows-amd64
	@$(MAKE) windows-arm64
	@echo "Release complete! Built binaries for:"
	@echo "  - bit2bin-linux-amd64"
	@echo "  - bit2bin-linux-arm64"
	@echo "  - bit2bin-windows-amd64.exe"
	@echo "  - bit2bin-windows-arm64.exe"
	@ls -lh $(BUILD_DIR)/

# Default target: build all platforms
all: linux-amd64 linux-arm64 windows-amd64 windows-arm64

# Linux AMD64
linux-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go

# Linux AArch64 (ARM64)
linux-arm64:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 main.go

# Windows AMD64
windows-amd64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go

# Windows AArch64 (ARM64)
windows-arm64:
	GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe main.go

# Clean build directory
clean:
	rm -rf $(BUILD_DIR)
	@echo "Cleaned build directory"
