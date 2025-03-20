package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestSuite is a struct to organize our test cases
type TestSuite struct {
	Name        string
	TestCases   []TestCase
	SetupScript string
}

// TestCase represents a single test case
type TestCase struct {
	Name           string
	Code           string
	ExpectedOutput string
	ShouldError    bool
	ErrorMessage   string
}

// Run the entire test suite
func (ts *TestSuite) Run(t *testing.T) {
	t.Run(ts.Name, func(t *testing.T) {
		// Create a temporary directory for test files
		tempDir, err := os.MkdirTemp("", "goception_test")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Run setup script if provided
		if ts.SetupScript != "" {
			setupFile := filepath.Join(tempDir, "setup.gct")
			if err := os.WriteFile(setupFile, []byte(ts.SetupScript), 0644); err != nil {
				t.Fatalf("Failed to write setup script: %v", err)
			}
			cmd := exec.Command("go", "run", "../main.go", setupFile)
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Setup script failed: %v", err)
			}
		}

		// Run each test case
		for _, tc := range ts.TestCases {
			t.Run(tc.Name, func(t *testing.T) {
				// Write test code to a temporary file
				testFile := filepath.Join(tempDir, tc.Name+".gct")
				if err := os.WriteFile(testFile, []byte(tc.Code), 0644); err != nil {
					t.Fatalf("Failed to write test file: %v", err)
				}

				// Run the test
				cmd := exec.Command("go", "run", "../main.go", testFile)
				output, _ := cmd.CombinedOutput() // Ignore execution error - we handle it later
				outputStr := string(output)

				// Check if we expect an error
				if tc.ShouldError {
					if !strings.Contains(outputStr, "ERROR") {
						t.Errorf("Expected error but got none. Output: %s", outputStr)
					} else if tc.ErrorMessage != "" && !strings.Contains(outputStr, tc.ErrorMessage) {
						t.Errorf("Expected error message to contain '%s', got: %s", tc.ErrorMessage, outputStr)
					}
				} else {
					// We don't expect an error
					if strings.Contains(outputStr, "ERROR") {
						t.Errorf("Unexpected error: %s", outputStr)
					} else {
						// Trim trailing null from output string that appears in Goception results
						trimmedOutput := strings.TrimSuffix(strings.TrimSpace(outputStr), "null")
						trimmedOutput = strings.TrimSpace(trimmedOutput)

						// Compare output with expected
						if tc.ExpectedOutput != "" && trimmedOutput != tc.ExpectedOutput {
							t.Errorf("Expected output: '%s', got: '%s'", tc.ExpectedOutput, trimmedOutput)
						}
					}
				}
			})
		}
	})
}

// TestVariablesAndConstants tests variable and constant declarations
func TestVariablesAndConstants(t *testing.T) {
	suite := TestSuite{
		Name: "VariablesAndConstants",
		TestCases: []TestCase{
			{
				Name: "SimpleVarDeclaration",
				Code: `
					var x: int = 5;
					print(x);
				`,
				ExpectedOutput: "5",
			},
			{
				Name: "SimpleConstDeclaration",
				Code: `
					const PI: int = 3;
					print(PI);
				`,
				ExpectedOutput: "3",
			},
			{
				Name: "MultipleDeclarations",
				Code: `
					var a: int = 1;
					var b: int = 2;
					var c: int = 3;
					print(a);
					print(b);
					print(c);
				`,
				ExpectedOutput: "1\n2\n3",
			},
			{
				Name: "VarReassignment",
				Code: `
					var count: int = 10;
					print(count);
					count = 20;
					print(count);
				`,
				ExpectedOutput: "10\n20",
			},
			{
				Name: "ConstReassignmentError",
				Code: `
					const MAX: int = 100;
					MAX = 200; // Should cause an error
				`,
				ShouldError:  true,
				ErrorMessage: "assignment to constant variable",
			},
			{
				Name: "TypeMismatchVarError",
				Code: `
					var age: int = "twenty"; // Should cause an error
				`,
				ShouldError:  true,
				ErrorMessage: "type mismatch",
			},
		},
	}
	suite.Run(t)
}

// TestStringOperations tests string operations
func TestStringOperations(t *testing.T) {
	suite := TestSuite{
		Name: "StringOperations",
		TestCases: []TestCase{
			{
				Name: "StringConcatenation",
				Code: `
					const greeting: string = "Hello, ";
					const name: string = "World";
					print(greeting + name + "!");
				`,
				ExpectedOutput: "Hello, World!",
			},
			{
				Name: "StringWithIntConcatenation",
				Code: `
					const message: string = "The answer is: ";
					const answer: int = 42;
					print(message + answer);
				`,
				ExpectedOutput: "The answer is: 42",
			},
			{
				Name: "StringWithBoolConcatenation",
				Code: `
					const prefix: string = "Is it true? ";
					const value: bool = true;
					print(prefix + value);
				`,
				ExpectedOutput: "Is it true? true",
			},
			{
				Name: "ComplexStringConcatenation",
				Code: `
					const name: string = "Alice";
					const age: int = 30;
					const isStudent: bool = false;
					print(name + " is " + age + " years old and student status is " + isStudent);
				`,
				ExpectedOutput: "Alice is 30 years old and student status is false",
			},
		},
	}
	suite.Run(t)
}

// TestArithmeticOperations tests arithmetic operations
func TestArithmeticOperations(t *testing.T) {
	suite := TestSuite{
		Name: "ArithmeticOperations",
		TestCases: []TestCase{
			{
				Name: "Addition",
				Code: `
					const a: int = 5;
					const b: int = 3;
					print(a + b);
				`,
				ExpectedOutput: "8",
			},
			{
				Name: "Subtraction",
				Code: `
					const a: int = 10;
					const b: int = 4;
					print(a - b);
				`,
				ExpectedOutput: "6",
			},
			{
				Name: "Multiplication",
				Code: `
					const a: int = 6;
					const b: int = 7;
					print(a * b);
				`,
				ExpectedOutput: "42",
			},
			{
				Name: "Division",
				Code: `
					const a: int = 20;
					const b: int = 5;
					print(a / b);
				`,
				ExpectedOutput: "4",
			},
			{
				Name: "ComplexExpression",
				Code: `
					print(2 + 3 * 4);
				`,
				ExpectedOutput: "14",
			},
			{
				Name: "GroupedExpression",
				Code: `
					print((2 + 3) * 4);
				`,
				ExpectedOutput: "20",
			},
			{
				Name: "NegativeNumbers",
				Code: `
					print(-5 + 3);
				`,
				ExpectedOutput: "-2",
			},
		},
	}
	suite.Run(t)
}

// TestComparisonOperations tests comparison operations
func TestComparisonOperations(t *testing.T) {
	suite := TestSuite{
		Name: "ComparisonOperations",
		TestCases: []TestCase{
			{
				Name: "Equals",
				Code: `
					print(5 == 5);
					print(5 == 6);
				`,
				ExpectedOutput: "true\nfalse",
			},
			{
				Name: "NotEquals",
				Code: `
					print(5 != 6);
					print(5 != 5);
				`,
				ExpectedOutput: "true\nfalse",
			},
			{
				Name: "LessThan",
				Code: `
					print(3 < 5);
					print(5 < 3);
					print(5 < 5);
				`,
				ExpectedOutput: "true\nfalse\nfalse",
			},
			{
				Name: "GreaterThan",
				Code: `
					print(7 > 3);
					print(3 > 7);
					print(3 > 3);
				`,
				ExpectedOutput: "true\nfalse\nfalse",
			},
			{
				Name: "LessThanOrEqual",
				Code: `
					print(3 <= 5);
					print(5 <= 5);
					print(7 <= 5);
				`,
				ExpectedOutput: "true\ntrue\nfalse",
			},
			{
				Name: "GreaterThanOrEqual",
				Code: `
					print(7 >= 5);
					print(5 >= 5);
					print(3 >= 5);
				`,
				ExpectedOutput: "true\ntrue\nfalse",
			},
		},
	}
	suite.Run(t)
}

// TestConditionals tests conditional operations
func TestConditionals(t *testing.T) {
	suite := TestSuite{
		Name: "Conditionals",
		TestCases: []TestCase{
			{
				Name: "IfTrue",
				Code: `
					if (true) {
						print("condition is true");
					}
				`,
				ExpectedOutput: "condition is true",
			},
			{
				Name: "IfFalse",
				Code: `
					if (false) {
						print("condition is true");
					} else {
						print("condition is false");
					}
				`,
				ExpectedOutput: "condition is false",
			},
			{
				Name: "IfWithComparisonTrue",
				Code: `
					const x: int = 10;
					if (x > 5) {
						print("x is greater than 5");
					}
				`,
				ExpectedOutput: "x is greater than 5",
			},
			{
				Name: "IfWithComparisonFalse",
				Code: `
					const x: int = 3;
					if (x > 5) {
						print("x is greater than 5");
					} else {
						print("x is not greater than 5");
					}
				`,
				ExpectedOutput: "x is not greater than 5",
			},
			{
				Name: "NestedIf",
				Code: `
					const x: int = 10;
					const y: int = 20;
					if (x > 5) {
						if (y > 15) {
							print("x > 5 and y > 15");
						} else {
							print("x > 5 but y <= 15");
						}
					} else {
						print("x <= 5");
					}
				`,
				ExpectedOutput: "x > 5 and y > 15",
			},
		},
	}
	suite.Run(t)
}

// TestFunctions tests function declarations and calls
func TestFunctions(t *testing.T) {
	suite := TestSuite{
		Name: "Functions",
		TestCases: []TestCase{
			{
				Name: "SimpleFunctionCall",
				Code: `
					const greet = function(): string {
						return "Hello, World!";
					};
					print(greet());
				`,
				ExpectedOutput: "Hello, World!",
			},
			{
				Name: "FunctionWithParameters",
				Code: `
					const add = function(a: int, b: int): int {
						return a + b;
					};
					print(add(3, 4));
				`,
				ExpectedOutput: "7",
			},
			{
				Name: "FunctionWithMultipleReturnStatements",
				Code: `
					const abs = function(n: int): int {
						if (n < 0) {
							return -n;
						}
						return n;
					};
					print(abs(-5));
					print(abs(5));
				`,
				ExpectedOutput: "5\n5",
			},
			{
				Name: "RecursiveFunction",
				Code: `
					const factorial = function(n: int): int {
						if (n <= 1) {
							return 1;
						} else {
							return n * factorial(n - 1);
						}
					};
					print(factorial(5));
				`,
				ExpectedOutput: "120",
			},
			{
				Name: "HigherOrderFunction",
				Code: `
					const apply = function(fn: function, x: int): int {
						return fn(x);
					};
					const double = function(x: int): int {
						return x * 2;
					};
					print(apply(double, 5));
				`,
				ExpectedOutput: "10",
			},
			{
				Name: "TypeMismatchParameterError",
				Code: `
					const greet = function(name: string): string {
						return "Hello, " + name + "!";
					};
					print(greet(123)); // Should cause an error
				`,
				ShouldError:  true,
				ErrorMessage: "type mismatch",
			},
			{
				Name: "TypeMismatchReturnError",
				Code: `
					const getName = function(): string {
						return 42; // Should cause an error
					};
					print(getName());
				`,
				ShouldError:  true,
				ErrorMessage: "return type mismatch",
			},
		},
	}
	suite.Run(t)
}

// TestTypeSystem tests type annotations and type checking
func TestTypeSystem(t *testing.T) {
	suite := TestSuite{
		Name: "TypeSystem",
		TestCases: []TestCase{
			{
				Name: "BasicTypeAnnotations",
				Code: `
					const i: int = 42;
					const s: string = "hello";
					const b: bool = true;
					print(i);
					print(s);
					print(b);
				`,
				ExpectedOutput: "42\nhello\ntrue",
			},
			{
				Name: "FunctionTypeAnnotation",
				Code: `
					const add: function = function(a: int, b: int): int {
						return a + b;
					};
					print(add(3, 4));
				`,
				ExpectedOutput: "7",
			},
			{
				Name: "TypeMismatchAssignment",
				Code: `
					const age: int = "twenty"; // Should cause an error
				`,
				ShouldError:  true,
				ErrorMessage: "type mismatch",
			},
			{
				Name: "TypeMismatchFunctionParameter",
				Code: `
					const increment = function(n: int): int {
						return n + 1;
					};
					print(increment("one")); // Should cause an error
				`,
				ShouldError:  true,
				ErrorMessage: "type mismatch",
			},
			{
				Name: "TypeMismatchFunctionReturn",
				Code: `
					const getName = function(): string {
						return 42; // Should cause an error
					};
					print(getName());
				`,
				ShouldError:  true,
				ErrorMessage: "return type mismatch",
			},
		},
	}
	suite.Run(t)
}

// TestIntegration tests more complex integrated examples
func TestIntegration(t *testing.T) {
	suite := TestSuite{
		Name: "Integration",
		TestCases: []TestCase{
			{
				Name: "FactorialExample",
				Code: `
					const factorial = function(n: int): int {
						if (n <= 1) {
							return 1;
						} else {
							return n * factorial(n - 1);
						}
					};
					var result: int = factorial(5);
					print("Factorial of 5 is: " + result);
				`,
				ExpectedOutput: "Factorial of 5 is: 120",
			},
			{
				Name: "StringManipulation",
				Code: `
					const greeting: string = "Hello";
					var message: string = greeting + ", World!";
					print(message);
					
					const count: int = 5;
					message = "The count is: " + count;
					print(message);
					
					const is_true: bool = true;
					message = "The boolean value is: " + is_true;
					print(message);
				`,
				ExpectedOutput: "Hello, World!\nThe count is: 5\nThe boolean value is: true",
			},
			{
				Name: "FunctionAsParameter",
				Code: `
					const apply = function(fn: function, x: int): int {
						return fn(x);
					};
					
					const double = function(x: int): int {
						return x * 2;
					};
					
					const square = function(x: int): int {
						return x * x;
					};
					
					print("Double of 7: " + apply(double, 7));
					print("Square of 5: " + apply(square, 5));
				`,
				ExpectedOutput: "Double of 7: 14\nSquare of 5: 25",
			},
		},
	}
	suite.Run(t)
}

// TestMainSuite runs all the test suites
func TestMainSuite(t *testing.T) {
	TestVariablesAndConstants(t)
	TestStringOperations(t)
	TestArithmeticOperations(t)
	TestComparisonOperations(t)
	TestConditionals(t)
	TestFunctions(t)
	TestTypeSystem(t)
	TestIntegration(t)
}

// For using 'go test'
// No main function is needed when running with 'go test'
