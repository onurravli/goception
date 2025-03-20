# Goception Language Support

This extension provides syntax highlighting and code snippets for the Goception programming language.

## Features

- Syntax highlighting for Goception (.gct) files
- Code snippets for common Goception patterns
- Basic language configuration

## Requirements

No special requirements or dependencies needed.

## Extension Settings

This extension contributes the following settings:

- None currently

## Known Issues

This is an early version of the extension. Please report any issues to the repository.

## Release Notes

### 1.0.0

Add import support.

### 0.0.1

Initial release of Goception language support

## Snippets

This extension provides several useful code snippets for Goception development:

| Snippet      | Description                  |      Prefix |
| ------------ | ---------------------------- | ----------: |
| Variable     | Declare a new variable       |       `var` |
| Constant     | Declare a new constant       |     `const` |
| Function     | Declare a new function       |  `function` |
| If Statement | Create an if statement       |        `if` |
| If-Else      | Create an if-else statement  |    `ifelse` |
| Print        | Print to console             |     `print` |
| Type         | Add a type annotation        |      `type` |
| Factorial    | Recursive factorial function | `factorial` |
| Concatenate  | String concatenation         |    `concat` |

## Installation

### From VS Code Marketplace (Coming Soon)

Search for "Goception" in the VS Code extension marketplace and click install.

### Manual Installation

1. Download or clone this repository
2. Copy the `extensions/vscode` folder to your VS Code extensions folder:
   - Windows: `%USERPROFILE%\.vscode\extensions`
   - macOS/Linux: `~/.vscode/extensions`
3. Restart VS Code

## Development

To contribute to this extension:

1. Clone the repository
2. Navigate to the extension directory: `cd extensions/vscode`
3. Install dependencies: `npm install`
4. Make your changes
5. Test the extension by pressing F5 in VS Code

## About Goception

Goception is a statically typed programming language with a clean, expressive syntax. It combines the simplicity of scripting languages with the safety of static type checking.

For more information about the Goception language, see the main [Goception documentation](../../GUIDE.md).

## License

This extension is released under the MIT License. See the LICENSE file for details.
