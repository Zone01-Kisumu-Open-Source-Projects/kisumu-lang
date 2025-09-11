package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/interpreter"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/parser"
)

// ReadFile reads a file and executes its content.
func ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	fmt.Printf("Reading from file: %s\n", filename)

	// Read the entire file content
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return err
	}

	// Parse and execute the content
	l := lexer.New(content, lexer.DefaultConfig)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		return fmt.Errorf("parsing failed")
	}

	// Evaluate the program
	interpreter := interpreter.New()
	result := interpreter.Eval(program)
	if result != nil {
		fmt.Println(result.Inspect())
	}

	return nil
}
