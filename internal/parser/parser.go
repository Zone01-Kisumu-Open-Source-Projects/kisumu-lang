package parser

import (
	"fmt"
	"strconv"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/ast"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

// Parser represents the parser for your language.
type Parser struct {
	l          *lexer.Lexer // Reference to the lexer
	currentTok token.Token  // Current token being processed
	peekTok    token.Token  // Next token (lookahead)
	errors     []string     // Errors encountered during parsing
}

// New creates a new parser instance.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Initialize the current and peek tokens
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances the current and peek tokens.
func (p *Parser) nextToken() {
	p.currentTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

// Errors returns a slice of errors encountered during parsing.
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseIntLiteral() ast.Node {
	// Check that the current token is an integer
	if p.currentTok.Type != token.INT {
		p.errors = append(p.errors, "expected INT token")
		return nil
	}

	// Parse the integer value
	value, err := strconv.Atoi(p.currentTok.Literal)
	if err != nil {
		p.errors = append(p.errors, "could not parse INT token")
		return nil
	}

	// Create and return the IntNode
	Intnode := &ast.IntNode{
		Token: p.currentTok,
		Value: value,
	}

	fmt.Println("IntNode =>> ", Intnode)
	return Intnode
}
