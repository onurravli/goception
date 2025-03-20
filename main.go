package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/onurravli/goception/evaluator"
	"github.com/onurravli/goception/lexer"
	"github.com/onurravli/goception/object"
	"github.com/onurravli/goception/parser"
)

func main() {
	if len(os.Args) > 1 {
		// If a file is provided, execute it
		filename := os.Args[1]
		executeFile(filename)
	} else {
		// Otherwise, start the REPL
		fmt.Println("Goception - A small and fast scripting language written in Go")
		fmt.Println("Type in commands")
		startRepl(os.Stdin, os.Stdout)
	}
}

func executeFile(filename string) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func startRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
