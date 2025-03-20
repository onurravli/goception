package test

import (
	"testing"
)

// TestGoception runs all the Goception language tests
func TestGoception(t *testing.T) {
	// Run each test suite
	TestVariablesAndConstants(t)
	TestStringOperations(t)
	TestArithmeticOperations(t)
	TestComparisonOperations(t)
	TestConditionals(t)
	TestFunctions(t)
	TestTypeSystem(t)
	TestIntegration(t)
}
