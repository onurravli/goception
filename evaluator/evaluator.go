package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/onurravli/goception/ast"
	"github.com/onurravli/goception/lexer"
	"github.com/onurravli/goception/object"
	"github.com/onurravli/goception/parser"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}

	// Track imported files to prevent circular imports
	importedFiles = make(map[string]bool)

	// Cache of imported file contents to avoid reading the same file multiple times
	importCache = make(map[string]string)
)

// Eval evaluates the given node and returns an object
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.VarStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		// Type checking if a type annotation is provided
		if node.Type != nil {
			if !checkType(val, node.Type.Value) {
				return newError("type mismatch: expected %s, got %s", node.Type.Value, val.Type())
			}
		}

		env.Set(node.Name.Value, val)
	case *ast.ConstStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		// Type checking if a type annotation is provided
		if node.Type != nil {
			if !checkType(val, node.Type.Value) {
				return newError("type mismatch: expected %s, got %s", node.Type.Value, val.Type())
			}
		}

		env.SetConst(node.Name.Value, val)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.ImportStatement:
		return evalImportStatement(node, env)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		returnType := node.ReturnType
		return &object.Function{
			Parameters: extractParameterNames(params),
			ParamTypes: extractParameterTypes(params),
			Body:       body,
			Env:        env,
			ReturnType: returnType,
		}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		// Check parameter types if function has type annotations
		if fn, ok := function.(*object.Function); ok && len(fn.ParamTypes) > 0 {
			for i, paramType := range fn.ParamTypes {
				if paramType != "" && i < len(args) {
					if !checkType(args[i], paramType) {
						return newError("type mismatch for argument %d: expected %s, got %s",
							i, paramType, args[i].Type())
					}
				}
			}
		}

		result := applyFunction(function, args)

		// Check return type if function has return type annotation
		if fn, ok := function.(*object.Function); ok && fn.ReturnType != nil {
			if !checkType(result, fn.ReturnType.Value) {
				return newError("return type mismatch: expected %s, got %s",
					fn.ReturnType.Value, result.Type())
			}
		}

		return result
	case *ast.AssignmentExpression:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		if !env.Reassign(node.Name.Value, val) {
			return newError("assignment to constant variable: %s", node.Name.Value)
		}

		return val
	}

	return NULL
}

// evalProgram evaluates a program
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

// evalBlockStatement evaluates a block statement
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

// evalPrefixExpression evaluates a prefix expression
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

// evalBangOperatorExpression evaluates a bang operator
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// evalMinusPrefixOperatorExpression evaluates a minus prefix operator
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// evalInfixExpression evaluates an infix expression
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && operator == "+":
		return evalStringConcatenation(left, right)
	case right.Type() == object.STRING_OBJ && operator == "+":
		return evalStringConcatenation(right, left)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

// evalStringConcatenation evaluates string concatenation with any type
func evalStringConcatenation(str object.Object, other object.Object) object.Object {
	stringVal := str.(*object.String).Value

	switch other.Type() {
	case object.INTEGER_OBJ:
		intVal := other.(*object.Integer).Value
		return &object.String{Value: stringVal + strconv.FormatInt(intVal, 10)}
	case object.BOOLEAN_OBJ:
		boolVal := other.(*object.Boolean).Value
		return &object.String{Value: stringVal + strconv.FormatBool(boolVal)}
	case object.NULL_OBJ:
		return &object.String{Value: stringVal + "null"}
	default:
		return &object.String{Value: stringVal + other.Inspect()}
	}
}

// evalIntegerInfixExpression evaluates an infix expression with integer operands
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

// evalStringInfixExpression evaluates an infix expression with string operands
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

// evalIfExpression evaluates an if expression
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

// isTruthy determines if an object is truthy
func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// evalIdentifier evaluates an identifier
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

// evalExpressions evaluates expressions
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

// applyFunction applies a function to arguments
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

// extendFunctionEnv extends the environment for a function
func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param, args[paramIdx])
		}
	}

	return env
}

// unwrapReturnValue unwraps a return value
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

// Helper functions

// nativeBoolToBooleanObject converts a native bool to a Boolean object
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// isError checks if an object is an error
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

// newError creates a new error
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// Builtins
var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}

// Helper function to extract parameter names from FunctionParameters
func extractParameterNames(params []*ast.FunctionParameter) []string {
	names := []string{}
	for _, param := range params {
		names = append(names, param.Name)
	}
	return names
}

// Helper function to extract parameter types from FunctionParameters
func extractParameterTypes(params []*ast.FunctionParameter) []string {
	types := []string{}
	for _, param := range params {
		if param.Type != nil {
			types = append(types, param.Type.Value)
		} else {
			types = append(types, "")
		}
	}
	return types
}

// checkType verifies if the object matches the expected type
func checkType(obj object.Object, typeName string) bool {
	switch typeName {
	case "int":
		return obj.Type() == object.INTEGER_OBJ
	case "string":
		return obj.Type() == object.STRING_OBJ
	case "bool":
		return obj.Type() == object.BOOLEAN_OBJ
	case "function":
		return obj.Type() == object.FUNCTION_OBJ
	default:
		return true // Unknown types are accepted for now
	}
}

// evalImportStatement imports and evaluates a file
func evalImportStatement(node *ast.ImportStatement, env *object.Environment) object.Object {
	filePath := node.Path.Value

	// Check if the file has a .gct extension, add if needed
	if !strings.HasSuffix(filePath, ".gct") {
		filePath = filePath + ".gct"
	}

	fmt.Printf("Importing file: %s\n", filePath)

	// Use absolute path for tracking imports
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return newError("could not resolve absolute path: %s", err.Error())
	}

	fmt.Printf("Absolute path: %s\n", absPath)

	// Check for circular imports
	if importedFiles[absPath] {
		// File is already being imported, skip to prevent circularity
		fmt.Printf("Skipping circular import of %s\n", filePath)
		return NULL
	}

	// Mark file as being imported
	importedFiles[absPath] = true
	defer func() {
		// After processing, unmark the file to allow it to be imported again in other contexts
		importedFiles[absPath] = false
	}()

	var input string
	var loadedFrom string

	// Check cache first
	if cachedInput, ok := importCache[absPath]; ok {
		input = cachedInput
		loadedFrom = "cache"
	} else {
		// First, try to read from the current directory
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			// If not found, try the examples directory
			examplesPath := filepath.Join("examples", filePath)
			fmt.Printf("Trying examples path: %s\n", examplesPath)
			fileBytes, err = os.ReadFile(examplesPath)
			if err != nil {
				// Try with just the basename in examples directory
				baseName := filepath.Base(filePath)
				examplesPath = filepath.Join("examples", baseName)
				fmt.Printf("Trying examples path with basename: %s\n", examplesPath)
				fileBytes, err = os.ReadFile(examplesPath)
				if err != nil {
					return newError("could not import file: %s. Tried: %s, %s, and %s",
						err.Error(), filePath, filepath.Join("examples", filePath), examplesPath)
				}
				loadedFrom = "examples with basename"
			} else {
				loadedFrom = "examples"
			}

			// Update absPath to the examples path
			absPath, _ = filepath.Abs(examplesPath)
		} else {
			loadedFrom = "current directory"
		}

		input = string(fileBytes)
		fmt.Printf("Loaded file from: %s, length: %d bytes\n", loadedFrom, len(input))

		// Cache the file content
		importCache[absPath] = input
	}

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		var errMsg strings.Builder
		errMsg.WriteString(fmt.Sprintf("parser errors in imported file %s (loaded from %s):\n", filePath, loadedFrom))
		for _, msg := range p.Errors() {
			errMsg.WriteString(fmt.Sprintf("\t%s\n", msg))
		}
		fmt.Printf("First few characters of file: %q\n", input[:min(20, len(input))])
		return newError(errMsg.String())
	}

	// Create new enclosed environment for the imported file
	importedEnv := object.NewEnclosedEnvironment(env)

	// Evaluate the imported program
	result := Eval(program, importedEnv)
	if isError(result) {
		return result
	}

	// Copy all variables from imported environment to the current environment
	importedEnv.ExportTo(env)

	fmt.Printf("Successfully imported: %s\n", filePath)
	return NULL
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
