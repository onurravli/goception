#!/bin/bash
# Simple installation script for Goception

echo "Goception Installer"
echo "==================="

# Set version
VERSION="1.0.0"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

# Build the binary
echo "Building Goception..."
go build -o goception main.go
if [ $? -ne 0 ]; then
    echo "Error: Build failed."
    exit 1
fi

# Ask where to install
echo ""
echo "Select installation option:"
echo "1) Install to /usr/local/bin (requires sudo)"
echo "2) Install to $HOME/bin"
echo "3) Install to current directory only"
read -p "Enter option (1-3): " INSTALL_OPTION

case $INSTALL_OPTION in
    1)
        INSTALL_PATH="/usr/local/bin"
        echo "Installing to $INSTALL_PATH..."
        sudo cp goception $INSTALL_PATH/
        if [ $? -ne 0 ]; then
            echo "Error: Installation failed. Do you have sudo privileges?"
            exit 1
        fi
        echo "Installation complete! You can now use 'goception' from anywhere."
        ;;
    2)
        INSTALL_PATH="$HOME/bin"
        echo "Installing to $INSTALL_PATH..."
        mkdir -p $INSTALL_PATH
        cp goception $INSTALL_PATH/
        
        # Check if $HOME/bin is in PATH
        if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
            echo "Adding $HOME/bin to PATH in your profile..."
            echo 'export PATH="$HOME/bin:$PATH"' >> $HOME/.profile
            echo "Please run 'source $HOME/.profile' or restart your terminal to update your PATH."
        fi
        echo "Installation complete! You can now use 'goception' command."
        ;;
    3)
        echo "Binary built in current directory."
        echo "You can run it with './goception'"
        ;;
    *)
        echo "Invalid option. Exiting."
        exit 1
        ;;
esac

# Ask if user wants to build VSCode extension
echo ""
read -p "Do you want to build the VSCode extension? (y/n): " BUILD_EXTENSION

if [ "$BUILD_EXTENSION" = "y" ] || [ "$BUILD_EXTENSION" = "Y" ]; then
    # Check if Node.js and npm are installed
    if ! command -v npm &> /dev/null; then
        echo "Error: npm is not installed. Please install Node.js and npm first."
        exit 1
    fi
    
    # Check if vsce is installed
    if ! command -v vsce &> /dev/null; then
        echo "Installing vsce..."
        npm install -g @vscode/vsce
    fi
    
    echo "Building VSCode extension..."
    cd extensions/vscode
    npm install
    vsce package
    
    echo "VSCode extension built successfully."
    echo "You can install it in VSCode by running:"
    echo "code --install-extension extensions/vscode/goception-vscode-*.vsix"
fi

echo ""
echo "Thank you for installing Goception!" 