/*
The REPL provides an environment / playground to interact with the language.

Although the language is not currently ready, the REPL will be used to visualize the lexer's functionality in the language of breaking source code to tokens.
*/
package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/interpreter"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/parser"
)

const PROMPT = ">>> "

// Start launches the REPL for language testing.
func Start() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Kisumu Lang REPL!")
	fmt.Println("Type your code below. Type 'exit' to quit.")
	fmt.Println("------------------------------------------")

	interpreter := interpreter.New()

	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			fmt.Println("\nGoodbye!")
			break
		}
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("Exiting REPL. Goodbye!")
			break
		}

		if input == "" {
			fmt.Println("Please enter some code or type 'exit' to quit.")
			continue
		}

		// Lex the input
		l := lexer.New(input, lexer.DefaultConfig)
		p := parser.New(l)

		program := p.ParseProgram()

		// Check for parsing errors
		if len(p.Errors()) != 0 {
			fmt.Println("Parser errors:")
			for _, err := range p.Errors() {
				fmt.Printf("  %s\n", err)
			}
			continue
		}

		// Evaluate the program
		result := interpreter.Eval(program)
		if result != nil {
			fmt.Println(result.Inspect())
		}
	}
	return nil
}
