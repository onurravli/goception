package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/onurravli/goception/ast"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
)

// Object represents an object in the VM
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represents an integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean represents a boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// ReturnValue represents a return value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error represents an error
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Function represents a function object
type Function struct {
	Parameters []string
	ParamTypes []string
	Body       *ast.BlockStatement
	Env        *Environment
	ReturnType *ast.TypeAnnotation
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for i, p := range f.Parameters {
		paramStr := p
		if i < len(f.ParamTypes) && f.ParamTypes[i] != "" {
			paramStr = p + ": " + f.ParamTypes[i]
		}
		params = append(params, paramStr)
	}

	out.WriteString("function")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")

	if f.ReturnType != nil {
		out.WriteString(": ")
		out.WriteString(f.ReturnType.Value)
	}

	out.WriteString(" {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String represents a string
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// BuiltinFunction represents a builtin function
type BuiltinFunction func(args ...Object) Object

// Builtin represents a builtin object
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Environment represents a variable environment
type Environment struct {
	store     map[string]Object
	outer     *Environment
	constants map[string]bool // Track which variables are constants
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	c := make(map[string]bool)
	return &Environment{store: s, constants: c}
}

// NewEnclosedEnvironment creates a new enclosed environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get gets a variable from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
		return obj, ok
	}
	return obj, ok
}

// Set sets a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	e.constants[name] = false
	return val
}

// SetConst sets a constant variable in the environment
func (e *Environment) SetConst(name string, val Object) Object {
	e.store[name] = val
	e.constants[name] = true
	return val
}

func (e *Environment) Reassign(name string, val Object) bool {
	if _, ok := e.store[name]; ok {
		// If variable exists in current environment
		if e.constants[name] {
			return false // Cannot reassign constants
		}
		e.store[name] = val
		return true
	}

	// Check if variable exists in outer environment
	if e.outer != nil {
		return e.outer.Reassign(name, val)
	}

	// Variable doesn't exist anywhere, treat as new variable
	e.store[name] = val
	e.constants[name] = false
	return true
}
