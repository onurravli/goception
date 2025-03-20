# Goception Test Suite

This directory contains a comprehensive test suite for the Goception programming language. The tests are designed to verify the correctness and robustness of the language implementation.

## Structure

The test suite is organized into the following test categories:

1. **Variables and Constants** - Tests variable and constant declarations, reassignments, and type checking
2. **String Operations** - Tests string concatenation and string operations
3. **Arithmetic Operations** - Tests arithmetic expressions and operators
4. **Comparison Operations** - Tests comparison operators and boolean results
5. **Conditionals** - Tests if/else statements and conditional logic
6. **Functions** - Tests function declarations, parameters, return values, and recursion
7. **Type System** - Tests type annotations and type checking
8. **Integration** - Tests complex examples combining multiple language features

## Running the Tests

You can run the test suite in two ways:

### Using Go's testing framework

```bash
cd test
go test -v
```

### Using the convenience script

```bash
cd test
chmod +x run_tests.sh
./run_tests.sh
```

## Adding New Tests

To add new tests, follow these steps:

1. Decide which test category your test belongs to or create a new category
2. Add a new test case to the appropriate TestSuite in `test_suite.go`
3. Follow the format of existing test cases, providing:
   - A unique name
   - The Goception code to test
   - The expected output or error message
   - Whether the test should result in an error

Example of adding a new test case:

```go
TestCase{
    Name: "MyNewTest",
    Code: `
        const x: int = 42;
        print(x);
    `,
    ExpectedOutput: "42",
},
```

Or for a test that should produce an error:

```go
TestCase{
    Name: "MyNewErrorTest",
    Code: `
        const x: int = "not a number";
    `,
    ShouldError: true,
    ErrorMessage: "type mismatch",
},
```

## Test Framework

The test framework uses Go's built-in testing package and provides a flexible way to define and run Goception code snippets. It automatically:

1. Creates temporary files for each test case
2. Runs the Goception interpreter on each file
3. Captures and verifies the output or error messages
4. Provides clear error reporting for failures

## Troubleshooting

If you encounter issues when running the tests:

1. Make sure the Goception interpreter is correctly built
2. Verify that the paths to the test files are correct
3. Check that the expected outputs exactly match the actual outputs (including whitespace)
4. For error cases, ensure that the error message contains the specified substring
