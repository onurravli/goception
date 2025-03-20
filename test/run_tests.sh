#!/bin/bash

# Simple test runner for Goception tests
echo "Running Goception test suite..."

# Navigate to the test directory
cd "$(dirname "$0")"

# Run the Go tests with the appropriate files
go test -v goception_test.go test_suite.go

# Check the exit code
if [ $? -eq 0 ]; then
    echo -e "\n\033[32mAll tests passed!\033[0m"
    exit 0
else
    echo -e "\n\033[31mTest suite failed!\033[0m"
    exit 1
fi 