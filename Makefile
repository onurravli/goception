.PHONY: all build install clean test macos-install

BINARY_NAME := goception
INSTALL_PATH := $(GOPATH)/bin
MACOS_INSTALL_PATH := /usr/local/bin
VERSION := 1.0.0
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

all: build

build:
	@echo "Building Goception..."
	@go build $(LDFLAGS) -o $(BINARY_NAME) main.go

install: build
	@echo "Installing Goception to $(INSTALL_PATH)..."
	@mkdir -p $(INSTALL_PATH)
	@cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installation complete. You can now use 'goception' command."

macos-install: build
	@echo "Installing Goception to $(MACOS_INSTALL_PATH)..."
	@sudo cp $(BINARY_NAME) $(MACOS_INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installation complete. You can now use 'goception' command globally."

local-install: build
	@echo "Installing Goception to ./bin/..."
	@mkdir -p ./bin
	@cp $(BINARY_NAME) ./bin/$(BINARY_NAME)
	@echo "Installation complete. Binary available at ./bin/$(BINARY_NAME)"

test:
	@echo "Running tests..."
	@go test ./...

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -rf ./bin

# Build the VSCode extension
vscode-ext:
	@echo "Building VSCode extension..."
	@cd extensions/vscode && npm install && vsce package
	@echo "VSCode extension built. VSIX file available in extensions/vscode directory."

# Install the VSCode extension locally
vscode-ext-install: vscode-ext
	@echo "Installing VSCode extension..."
	@code --install-extension extensions/vscode/goception-vscode-*.vsix

# Create a release package including the compiler and extension
release: build vscode-ext
	@echo "Creating release package..."
	@mkdir -p release/$(BINARY_NAME)-$(VERSION)
	@cp $(BINARY_NAME) release/$(BINARY_NAME)-$(VERSION)/
	@cp -r examples release/$(BINARY_NAME)-$(VERSION)/
	@cp README.md release/$(BINARY_NAME)-$(VERSION)/
	@cp extensions/vscode/goception-vscode-*.vsix release/$(BINARY_NAME)-$(VERSION)/
	@cd release && tar -czvf $(BINARY_NAME)-$(VERSION).tar.gz $(BINARY_NAME)-$(VERSION)
	@echo "Release package created at release/$(BINARY_NAME)-$(VERSION).tar.gz"

help:
	@echo "Goception Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build         - Build the Goception compiler"
	@echo "  make install       - Install Goception to GOPATH/bin"
	@echo "  make macos-install - Install Goception to /usr/local/bin (requires sudo)"
	@echo "  make local-install - Install Goception to ./bin directory"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Remove binary and bin directory"
	@echo "  make vscode-ext    - Build the VSCode extension"
	@echo "  make vscode-ext-install - Install the VSCode extension locally"
	@echo "  make release       - Create a release package"
	@echo "  make help          - Display this help message" 