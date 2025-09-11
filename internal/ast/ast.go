package ast

import (
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

type Node interface {
	TokenLiteral() string // Returns the literal value of the token (e.g., "42").
	String() string       // Returns a string representation of the node.
}

type Expression interface {
	Node
	expressionNode()
}

// A binary expression node (e.g., `3 + 5`)
type BinaryExpression struct {
	Left     Expression // The left operand (e.g., an IntNode for `3`).
	Operator string     // The operator (e.g., `+`).
	Right    Expression // The right operand (e.g., an IntNode for `5`).
}

// Implement methods for BinaryExpression
func (b *BinaryExpression) TokenLiteral() string {
	return b.Operator
}

func (b *BinaryExpression) String() string {
	return "(" + b.Left.String() + " " + b.Operator + " " + b.Right.String() + ")"
}

func (b *BinaryExpression) expressionNode() {}

// StringNode represents a string literal
type StringNode struct {
	Token token.Token
	Value string
}

func (s *StringNode) TokenLiteral() string {
	return s.Token.Literal
}

func (s *StringNode) String() string {
	return s.Token.Literal
}

func (s *StringNode) expressionNode() {}

// BoolNode represents a boolean literal
type BoolNode struct {
	Token token.Token
	Value bool
}

func (b *BoolNode) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BoolNode) String() string {
	return b.Token.Literal
}

func (b *BoolNode) expressionNode() {}

// Identifier represents an identifier expression
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// Statement interface for statements
type Statement interface {
	Node
	statementNode()
}

// LetStatement represents a variable declaration
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out string
	out += ls.TokenLiteral() + " "
	out += ls.Name.String()
	out += " = "
	if ls.Value != nil {
		out += ls.Value.String()
	}
	out += ";"
	return out
}

func (ls *LetStatement) statementNode() {}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out string
	out += rs.TokenLiteral() + " "
	if rs.ReturnValue != nil {
		out += rs.ReturnValue.String()
	}
	out += ";"
	return out
}

func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}

// Program represents the root node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

// PrefixExpression represents a prefix expression (e.g., -5, !true)
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	return "(" + pe.Operator + pe.Right.String() + ")"
}

func (pe *PrefixExpression) expressionNode() {}

// InfixExpression represents an infix expression (e.g., 5 + 5)
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

func (ie *InfixExpression) expressionNode() {}

// IntNode represents an integer literal in the syntax tree
type IntNode struct {
	Token token.Token // represents the integer(e.g, '42';)
	Value int         // actual integer val an an int
}

// implements the Node interface
func (i *IntNode) TokenLiteral() string {
	return i.Token.Literal
}

// String implements the Node interface, returning a string representation of the node.
func (i *IntNode) String() string {
	return i.Token.Literal
}

// expressionNode implements the Expression interface
func (i *IntNode) expressionNode() {}
