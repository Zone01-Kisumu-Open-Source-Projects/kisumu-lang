package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

// ReadFile reads a file and tokenizes its content.
func ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return err
	}
	defer file.Close()

	fmt.Printf("Reading from file: %s\n", filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line, lexer.DefaultConfig)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
	return nil
}
