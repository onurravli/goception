## Basic Syntax

- Goception is a statically typed language, which means that the type of a variable is known at compile time.

```gct
const a: int = 23;
var b: string = "Hello, world";
```

- In Goception, we can define variables and constants with the `var` and `const` keywords. We can write comments with the `//` operator for single line comments and `/* ... */` for multi-line comments.

```gct
// This is a single line comment
/*
This is a multi-line comment
*/
```

- In Goception, we have to use `;` to end a statement.

```gct
var a: int = 23; // Correct
var a: int = 23 // Incorrect
```

- In Goception, constants written in uppercase are immutable, and variables written in lowercase are mutable with under-score notation.

```gct
const A_CONSTANT_VALUE: int = 23; // Correct
var a_variable_value: int = 23; // Correct, again, _but not recommended_
```

1. Variables and Constants, Types and Literals

In Goception, we have some primitive types:

- `int`: Represents integer numbers.

```gct
var x: int = 10; // Mutable
// or
const x: int = 10; // Immutable
```

- `float`: Represents floating-point numbers.

```gct
var y: float = 10.5; // Mutable
// or
const y: float = 10.5; // Immutable
```

- `bool`: Represents boolean values.

```gct
var z: bool = true; // Mutable
// or
const z: bool = true; // Immutable
```

- `char`: Represents a single character.

```gct
var c: char = 'a'; // Mutable
// or
const c: char = 'a'; // Immutable
```

- `string`: Represents text strings.

```gct
var s: string = "Hello, World!"; // Mutable
// or
const s: string = "Hello, World!"; // Immutable
```

- `array`: Represents ordered lists of other types.

```gct
var a: array<int> = [1, 2, 3]; // Mutable
// or
const a: array<int> = [1, 2, 3]; // Immutable
```

- `map`: Represents a collection of key-value pairs.

```gct
var m: map<string, int> = {"key1": 1, "key2": 2, "key3": 3}; // Mutable
// To get the value of a key, we can use the `[]` operator.
const value: int = m["key1"]; // We don't have to use `int` here, it will be inferred:
                              // const value: int = m["key1"];
// or
const m: map<string, int> = {"key1": 1, "key2": 2, "key3": 3}; // Immutable
```

In addition to these primitive types, Goception also has some compound types:

- `struct`: Represents a collection of named fields, like JSON objects.

```gct
var person: struct {
    name: string;
    age: int;
} = {name: "John", age: 30}; // Mutable
// And we can get the fields like this:
const name: string = person.name;
const age: int = person.age;
// or
const person: struct {name: string; age: int} = {name: "John", age: 30}; // Immutable
```

- `function`: Represents a function that can be called with arguments.

```gct
var add: function(a: int, b: int): int = {
    var sum: int = a + b;
    return sum;
};

// And we can call it like this:
const result: int = add(1, 2);
```

- `void`: Represents a function that does not return anything.

```gct
function print_hello_world(): void {
    print("Hello, World!");
}
```

2. Operators

- Arithmetic operators: `+`, `-`, `*`, `/`, `%`
- Comparison operators: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Logical operators: `&&`, `||`, `!`

3. Built-in Functions

- `print`: Prints a message to the console.

```gct
print("Hello, World!");
const name: string = "John";
print("My name is %s", name);
var age: int = 23;
age = age + 1;
print("I am %d years old", age);
```

- `typeof`: Returns the type of a variable.

```gct
const type_of_x: string = typeof(x);
```

- `len`: Returns the length of a string or array.

```gct
const length: int = len("Hello, World!");
// or
const length: int = len([1, 2, 3]);
```

- `append`: Appends an element to an array.

```gct
const arr: array<int> = [1, 2, 3];
const new_arr: array<int> = append(arr, 4);
```

- `pop`: Removes and returns the last element of an array.

```gct
const arr: array<int> = [1, 2, 3];
const last_element: int = pop(arr);
```

- `push`: Adds an element to the end of an array.

```gct
const arr: array<int> = [1, 2, 3];
const new_arr: array<int> = push(arr, 4);
```

3. File structure

Every Goception program, have to start with the `main` function. If there is no `main` function, the program will not run, and throws an error.

```gct

// main.gct

function main(): void {
    print("Hello, World!");
}

// other_file.gct

function other_function(): void {
    print("Hello, World!");
}

```

4. Importing files

We can import files with the `import` keyword.

```gct
// other_file.gct

function sum(a: int, b: int): int {
  var sum: int = a + b;
  return sum;
}
```

```gct
// main.gct

import "other_file.gct";

function main(): void {
    const result: int = sum(1, 2);
    print(result);
}
```

5. Control flow

- `if`: Executes a block of code if a condition is true.

```gct
var x: int = 10;

if (x > 0) {
    print("x is greater than 0");
}
```

- `else`: Executes a block of code if a condition is false.

```gct
var x: int = 10;


if (x > 0) {
    print("x is greater than 0");
} else {
    print("x is less than or equal to 0");
}
```

- `elseif`: Executes a block of code if a condition is true.

```gct
var x: int = 10;

if (x > 0) {
    print("x is greater than 0");
} elseif (x == 0) {
    print("x is equal to 0");
} else {
    print("x is less than 0");
}
```

- `while`: Executes a block of code while a condition is true.

```gct
var x: int = 0;

while (x < 10) {
    print(x);
    x = x + 1;
}
```

- `for`: Executes a block of code a specific number of times.

```gct
for (var i: int = 0; i < 10; i = i + 1) {
    print(i);
}
# or
var i: int = 0;

for (i = 0; i < 10; i = i + 1) {
    print(i);
}
```

- `break`: Breaks out of a loop.

```gct
for (var i: int = 0; i < 10; i = i + 1) {
    if (i == 5) {
        break;
    }
    print(i);
}
```

- `continue`: Skips the current iteration of a loop.

```gct
for (var i: int = 0; i < 10; i = i + 1) {
    if (i == 5) {
        continue;
    }
    print(i);
}
```

- `return`: Returns a value from a function.

```gct
function sum(a: int, b: int): int {
    var sum: int = a + b;
    return sum;
}

const result: int = sum(1, 2);
```
