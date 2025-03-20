# Goception

Goception is a small and fast scripting language written in Go.

## Installation

### Requirements

- Go 1.18 or higher
- For VSCode extension: Node.js and npm

### From Source

Clone the repository and build from source:

```bash
git clone https://github.com/onurravli/goception.git
cd goception
```

#### Build and Install

To build and install Goception, use the provided Makefile.

```bash
# Just build the binary
make build

# Install to GOPATH/bin
make install

# Install globally on macOS (requires sudo)
make macos-install

# Install locally in the project's bin directory
make local-install
```

#### VSCode Extension

To build and install the VSCode extension:

```bash
# Build the extension
make vscode-ext

# Install the extension in VSCode
make vscode-ext-install
```

## Usage

### Running Scripts

```bash
goception examples/factorial.gct
```

### Interactive Mode

```bash
goception
```

## Language Features

- Dynamic typing with optional type annotations
- First-class functions
- Variable and constant declarations
- Control flow statements (if/else)
- Module system with imports
- String concatenation with automatic type conversion
- Lexical scoping
- Recursive functions
- Comments (single-line and multi-line)

## Example

```
// Import utilities
import "math-utils.gct";

// Use imported functions and constants
print("Square of 5: " + square(5));
print("PI value: " + PI);

// Define a new function
const greet = function(name) {
  return "Hello, " + name + "!";
};

print(greet("Goception"));
```

## License

MIT
