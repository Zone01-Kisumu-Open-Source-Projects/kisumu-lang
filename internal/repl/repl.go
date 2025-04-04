/*
The REPL provides an environment / playground to interact with the language.

Although the language is not currently ready, the REPL will be used to visualize the lexer's functionality in the language of breaking source code to tokens.
*/
package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

const PROMPT = ">>> "

// Start launches the REPL for lexer testing.
func Start() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Kisumu Lang Lexer REPL!")
	fmt.Println("Type your code below. Type 'exit' to quit.")
	fmt.Println("------------------------------------------")

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

		l := lexer.New(input, lexer.DefaultConfig)
		fmt.Println("Tokens:")
		fmt.Println("------------------------------------------")

		// Print tokens until EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
		fmt.Println("------------------------------------------")
	}
	return nil
}
