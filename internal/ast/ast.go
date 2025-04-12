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

// Implement methods for BinaryExpression...

// Intnode represents an integer literal in the syntax tree
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
