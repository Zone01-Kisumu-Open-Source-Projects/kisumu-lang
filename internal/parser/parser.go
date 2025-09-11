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
	l              *lexer.Lexer // Reference to the lexer
	currentTok     token.Token  // Current token being processed
	peekTok        token.Token  // Next token (lookahead)
	errors         []string     // Errors encountered during parsing
	prefixParseFns map[string]prefixParseFn
	infixParseFns  map[string]infixParseFn
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

// ParseProgram parses the entire program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentTok.Type != token.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}

	return program
}

// parseStatement parses a statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentTok.Type {
	case token.CONST:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.FOR:
		return p.parseForStatement()
	case token.BREAK:
		return p.parseBreakStatement()
	case token.CONTINUE:
		return p.parseContinueStatement()
	case token.LEFTBRACE:
		return p.parseBlockStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement parses a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentTok}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentTok, Value: p.currentTok.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement parses a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentTok}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement parses an expression statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentTok}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Precedence constants
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Precedence map
var precedences = map[string]int{
	token.EQUAL:           EQUALS,
	token.NOTEQUAL:        EQUALS,
	token.LESSTHAN:        LESSGREATER,
	token.GREATERTHAN:     LESSGREATER,
	token.PLUS:            SUM,
	token.MINUS:           SUM,
	token.SLASH:           PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.LEFTPARENTHESIS: CALL,
}

// parseExpression parses an expression with given precedence
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentTok.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentTok.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekTok.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// parseIntegerLiteral parses an integer literal
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntNode{Token: p.currentTok}

	literal := p.currentTok.Literal
	var value int
	var err error

	// Handle different integer bases
	if len(literal) > 2 {
		switch literal[:2] {
		case "0x", "0X":
			val, parseErr := strconv.ParseInt(literal[2:], 16, 64)
			value, err = int(val), parseErr
		case "0b", "0B":
			val, parseErr := strconv.ParseInt(literal[2:], 2, 64)
			value, err = int(val), parseErr
		default:
			value, err = strconv.Atoi(literal)
		}
	} else {
		value, err = strconv.Atoi(literal)
	}

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = int(value)
	return lit
}

// parseStringLiteral parses a string literal
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringNode{Token: p.currentTok, Value: p.currentTok.Literal}
}

// parseBoolean parses a boolean literal
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.BoolNode{Token: p.currentTok, Value: p.currentTok.Type == token.TRUE}
}

// parseIdentifier parses an identifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentTok, Value: p.currentTok.Literal}
}

// parsePrefixExpression parses a prefix expression
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currentTok,
		Operator: p.currentTok.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// parseInfixExpression parses an infix expression
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentTok,
		Operator: p.currentTok.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

// parseGroupedExpression parses a grouped expression
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	return exp
}

// Helper methods
func (p *Parser) peekTokenIs(t string) bool {
	return p.peekTok.Type == t
}

func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t string) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekTok.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t string) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekTok.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currentTok.Type]; ok {
		return p
	}
	return LOWEST
}

// Parser state
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// New creates a new parser instance with parse functions
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[string]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.LEFTPARENTHESIS, p.parseGroupedExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	p.infixParseFns = make(map[string]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.NOTEQUAL, p.parseInfixExpression)
	p.registerInfix(token.LESSTHAN, p.parseInfixExpression)
	p.registerInfix(token.GREATERTHAN, p.parseInfixExpression)
	p.registerInfix(token.LEFTPARENTHESIS, p.parseCallExpression)

	// Initialize the current and peek tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefix(tokenType string, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType string, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// parseIfStatement parses an if statement
func (p *Parser) parseIfStatement() *ast.IfExpression {
	stmt := &ast.IfExpression{Token: p.currentTok}

	if !p.expectPeek(token.LEFTPARENTHESIS) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	if !p.expectPeek(token.LEFTBRACE) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LEFTBRACE) {
			return nil
		}

		stmt.Alternative = p.parseBlockStatement()
	}

	return stmt
}

// parseWhileStatement parses a while statement
func (p *Parser) parseWhileStatement() *ast.WhileLoop {
	stmt := &ast.WhileLoop{Token: p.currentTok}

	if !p.expectPeek(token.LEFTPARENTHESIS) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	if !p.expectPeek(token.LEFTBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseForStatement parses a for statement
func (p *Parser) parseForStatement() *ast.ForLoop {
	stmt := &ast.ForLoop{Token: p.currentTok}

	if !p.expectPeek(token.LEFTPARENTHESIS) {
		return nil
	}

	p.nextToken()

	// Parse initializer (optional)
	if p.currentTok.Type != token.SEMICOLON {
		if p.currentTok.Type == token.CONST {
			stmt.Initializer = p.parseLetStatement()
		} else {
			stmt.Initializer = p.parseExpressionStatement()
		}
		// Consume semicolon after initializer
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	p.nextToken()

	// Parse condition (optional)
	if p.currentTok.Type != token.SEMICOLON {
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	p.nextToken()

	// Parse update (optional)
	if p.currentTok.Type != token.RIGHTPARENTHESIS {
		if p.currentTok.Type == token.CONST {
			stmt.Update = p.parseLetStatement()
		} else {
			stmt.Update = p.parseExpressionStatement()
		}
		// Consume semicolon after update
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
	}

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	if !p.expectPeek(token.LEFTBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseBreakStatement parses a break statement
func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	stmt := &ast.BreakStatement{Token: p.currentTok}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseContinueStatement parses a continue statement
func (p *Parser) parseContinueStatement() *ast.ContinueStatement {
	stmt := &ast.ContinueStatement{Token: p.currentTok}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseBlockStatement parses a block statement
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currentTok}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenIs(token.RIGHTBRACE) && !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		block.Statements = append(block.Statements, stmt)
		p.nextToken()
	}

	return block
}

// parseFunctionLiteral parses a function literal
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.currentTok}

	if !p.expectPeek(token.LEFTPARENTHESIS) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LEFTBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseFunctionParameters parses function parameters
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RIGHTPARENTHESIS) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.currentTok, Value: p.currentTok.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.currentTok, Value: p.currentTok.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	return identifiers
}

// parseCallExpression parses a call expression
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.currentTok, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

// parseCallArguments parses call arguments
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RIGHTPARENTHESIS) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RIGHTPARENTHESIS) {
		return nil
	}

	return args
}

// Helper methods
func (p *Parser) currentTokenIs(t string) bool {
	return p.currentTok.Type == t
}
