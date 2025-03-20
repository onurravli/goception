package main

import (
	"fmt"
	"os"

	"github.com/onurravli/goception/lexer"
	"github.com/onurravli/goception/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug-lexer <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Read the file
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("File size: %d bytes\n", len(input))

	// Initialize the lexer
	l := lexer.New(string(input))

	// Tokenize and print each token
	var tokenCount int
	for {
		tok := l.NextToken()
		tokenCount++

		fmt.Printf("%4d | Type: %-10s | Literal: %-15q | Line: %d | Column: %d\n",
			tokenCount, tok.Type, tok.Literal, tok.Line, tok.Column)

		if tok.Type == token.EOF {
			break
		}

		// Emergency break to prevent infinite loops
		if tokenCount > 1000 {
			fmt.Println("Warning: Reached token limit (1000)")
			break
		}
	}

	fmt.Printf("Total tokens: %d\n", tokenCount)
}
