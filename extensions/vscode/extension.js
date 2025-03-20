// The module 'vscode' contains the VS Code extensibility API
const vscode = require("vscode");

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {
  console.log("Goception language extension is now active!");

  // Register a command to show information about Goception
  let disposable = vscode.commands.registerCommand(
    "goception.about",
    function () {
      vscode.window.showInformationMessage(
        "Goception Language Support is active. Happy coding!"
      );
    }
  );

  context.subscriptions.push(disposable);
}

// This method is called when your extension is deactivated
function deactivate() {}

module.exports = {
  activate,
  deactivate,
};
