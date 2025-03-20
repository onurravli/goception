# Goception Programming Language Guide

## Introduction

Goception is a statically typed programming language with a clean, expressive syntax inspired by languages like JavaScript and Go. It was developed to combine the simplicity of scripting languages with the safety of static type checking.

This guide covers the syntax, features, and usage of Goception to help you get started with the language.

## Table of Contents

- [Basic Syntax](#basic-syntax)
- [Variables and Constants](#variables-and-constants)
- [Data Types](#data-types)
- [Operators](#operators)
- [Control Flow](#control-flow)
- [Functions](#functions)
- [Type System](#type-system)
- [Built-in Functions](#built-in-functions)
- [Examples](#examples)
- [Testing](#testing)

## Basic Syntax

Goception uses semicolons to terminate statements and employs curly braces for code blocks.

```gct
// This is a single-line comment
/* This is a
   multi-line comment */

var x: int = 10;  // Variable declaration with type annotation
print(x);         // Function call
```

## Variables and Constants

Goception supports two types of variable declarations:

- `var` for variables that can be reassigned
- `const` for variables that cannot be reassigned

```gct
// Variable declaration
var count: int = 5;
count = 10;  // Valid: variables can be reassigned

// Constant declaration
const PI: int = 3;
// PI = 4;  // Invalid: constants cannot be reassigned
```

### Naming Conventions

Constants are typically written in uppercase, while variables use lowercase with underscores for readability:

```gct
const MAX_USERS: int = 100;  // Constant
var user_count: int = 0;     // Variable
```

## Data Types

Goception supports the following primitive data types:

### Integer (`int`)

Represents whole numbers.

```gct
var age: int = 30;
```

### String (`string`)

Represents text enclosed in double quotes.

```gct
var name: string = "Goception";
```

### Boolean (`bool`)

Represents true or false values.

```gct
var is_valid: bool = true;
var has_errors: bool = false;
```

### Function (`function`)

Represents a callable code block.

```gct
const add: function = function(a: int, b: int): int {
  return a + b;
};
```

### Null

Represents the absence of a value.

```gct
var result = null;
```

## Operators

Goception supports various operators for arithmetic, comparison, and logical operations.

### Arithmetic Operators

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Negation: `-` (unary)

```gct
var a: int = 5 + 3;  // 8
var b: int = 10 - 4; // 6
var c: int = 3 * 4;  // 12
var d: int = 8 / 2;  // 4
var e: int = -5;     // -5
```

### Comparison Operators

- Equal to: `==`
- Not equal to: `!=`
- Greater than: `>`
- Less than: `<`
- Greater than or equal to: `>=`
- Less than or equal to: `<=`

```gct
var is_equal: bool = (5 == 5);       // true
var is_not_equal: bool = (5 != 3);   // true
var is_greater: bool = (10 > 5);     // true
var is_less: bool = (3 < 7);         // true
var is_greater_equal: bool = (5 >= 5); // true
var is_less_equal: bool = (3 <= 3);   // true
```

### Logical Operators

- Logical NOT: `!`
- Logical AND: `&&` (not implemented yet)
- Logical OR: `||` (not implemented yet)

```gct
var is_valid: bool = true;
var is_not_valid: bool = !is_valid;  // false
```

### String Concatenation

The `+` operator is also used for string concatenation, with automatic type conversion:

```gct
var message: string = "Hello, " + "World!";  // "Hello, World!"
var count: int = 5;
var status: string = "Count: " + count;  // "Count: 5"
var is_active: bool = true;
var state: string = "Active: " + is_active;  // "Active: true"
```

## Control Flow

Goception provides standard control flow constructs to manage program execution.

### Conditionals

The `if`, `else` expressions allow conditional execution:

```gct
var x: int = 10;

if (x > 5) {
  print("x is greater than 5");
} else {
  print("x is less than or equal to 5");
}
```

### Functions and Returns

Functions can be defined with the `function` keyword and return values using the `return` statement:

```gct
const factorial = function(n: int): int {
  if (n <= 1) {
    return 1;
  } else {
    return n * factorial(n - 1);
  }
};

print(factorial(5));  // 120
```

## Functions

In Goception, functions are first-class citizens, which means they can be:

- Assigned to variables
- Passed as arguments to other functions
- Returned from functions

### Function Declaration

Functions are typically defined with the `function` keyword:

```gct
const add = function(a: int, b: int): int {
  return a + b;
};
```

### Function Parameters

Functions can have typed parameters:

```gct
const greet = function(name: string, age: int): string {
  return "Hello, " + name + "! You are " + age + " years old.";
};
```

### Function Return Types

Functions can have an explicit return type annotation:

```gct
const square = function(x: int): int {
  return x * x;
};
```

### Higher-Order Functions

Functions can take other functions as parameters:

```gct
const apply = function(fn: function, x: int): int {
  return fn(x);
};

const double = function(x: int): int {
  return x * 2;
};

print(apply(double, 5));  // 10
```

## Type System

Goception features a static type system with type annotations.

### Type Annotations

Type annotations are specified after variable names, parameters, and function return types:

```gct
var count: int = 10;
const message: string = "Hello";
const calculate = function(x: int, y: int): int {
  return x + y;
};
```

### Type Checking

Goception performs type checking at runtime:

1. When assigning values to variables or constants with type annotations
2. When passing arguments to functions with typed parameters
3. When returning values from functions with return type annotations

For example, this code would produce a type error:

```gct
const age: int = "thirty";  // Error: type mismatch: expected int, got STRING
```

### Type Inference

Currently, Goception requires explicit type annotations, but the actual type checking happens at runtime.

## Built-in Functions

Goception provides several built-in functions for common operations:

### `print()`

Outputs the value of its argument to the console.

```gct
print("Hello, World!");  // Outputs: Hello, World!
print(42);              // Outputs: 42
print(true);           // Outputs: true
```

### `len()`

Returns the length of a string.

```gct
var name: string = "Goception";
print(len(name));  // Outputs: 9
```

## Examples

Here are some complete examples to demonstrate Goception's features:

### Factorial Calculation

```gct
const factorial = function(n: int): int {
  if (n <= 1) {
    return 1;
  } else {
    return n * factorial(n - 1);
  }
};

var result: int = factorial(5);
print("Factorial of 5 is: " + result);  // Outputs: Factorial of 5 is: 120
```

### String Manipulation

```gct
const greeting: string = "Hello";
var message: string = greeting + ", World!";
print(message);  // Outputs: Hello, World!

// Concatenation with other types
const count: int = 5;
message = "The count is: " + count;
print(message);  // Outputs: The count is: 5

// Concatenation with boolean
const is_true: bool = true;
message = "The boolean value is: " + is_true;
print(message);  // Outputs: The boolean value is: true
```

### Function as Parameter

```gct
const apply = function(fn: function, x: int): int {
  return fn(x);
};

const double = function(x: int): int {
  return x * 2;
};

const square = function(x: int): int {
  return x * x;
};

print("Double of 7: " + apply(double, 7));  // Outputs: Double of 7: 14
print("Square of 5: " + apply(square, 5));  // Outputs: Square of 5: 25
```

## Testing

Goception comes with a comprehensive test suite to ensure the language implementation is robust and reliable. The tests cover various aspects of the language from basic syntax to complex functionality.

### Test Suite Structure

The test suite is organized into categories:

1. **Variables and Constants** - Tests for variable and constant declarations
2. **String Operations** - Tests for string concatenation and manipulation
3. **Arithmetic Operations** - Tests for arithmetic expressions
4. **Comparison Operations** - Tests for comparison operators
5. **Conditionals** - Tests for if/else statements
6. **Functions** - Tests for function declarations and calls
7. **Type System** - Tests for type annotations and checking
8. **Integration** - Tests for combining multiple language features

### Running the Tests

To run the tests, use one of the following methods:

```bash
# Using Go's testing framework
cd test
go test -v

# Or using the convenience script
cd test
chmod +x run_tests.sh
./run_tests.sh
```

### Writing Your Own Tests

You can extend the test suite with your own tests by following the pattern established in the `test/test_suite.go` file. See the `test/README.md` file for detailed instructions on adding new tests.

## Conclusion

Goception is a simple yet powerful programming language that combines the flexibility of scripting languages with the safety of static typing. While it's still under development, it already offers a solid foundation for writing clean and maintainable code.

As you become more familiar with Goception, you'll appreciate its straightforward syntax and powerful features. Happy coding!
