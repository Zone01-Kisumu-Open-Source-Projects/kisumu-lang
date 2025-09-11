package ast

import (
	"strings"

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

// IfExpression represents an if-else expression
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out string
	out += "if (" + ie.Condition.String() + ") " + ie.Consequence.String()
	if ie.Alternative != nil {
		out += " else " + ie.Alternative.String()
	}
	return out
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) statementNode()  {}

// BlockStatement represents a block of statements
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out string
	for _, s := range bs.Statements {
		out += s.String()
	}
	return out
}

func (bs *BlockStatement) statementNode() {}

// WhileLoop represents a while loop
type WhileLoop struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (wl *WhileLoop) TokenLiteral() string {
	return wl.Token.Literal
}

func (wl *WhileLoop) String() string {
	return "while (" + wl.Condition.String() + ") " + wl.Body.String()
}

func (wl *WhileLoop) statementNode() {}

// ForLoop represents a for loop
type ForLoop struct {
	Token       token.Token
	Initializer Statement
	Condition   Expression
	Update      Statement
	Body        *BlockStatement
}

func (fl *ForLoop) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *ForLoop) String() string {
	var out string
	out += "for ("
	if fl.Initializer != nil {
		out += fl.Initializer.String()
	}
	out += "; "
	if fl.Condition != nil {
		out += fl.Condition.String()
	}
	out += "; "
	if fl.Update != nil {
		out += fl.Update.String()
	}
	out += ") " + fl.Body.String()
	return out
}

func (fl *ForLoop) statementNode() {}

// FunctionLiteral represents a function literal
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out string
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out += "func(" + strings.Join(params, ", ") + ") " + fl.Body.String()
	return out
}

func (fl *FunctionLiteral) expressionNode() {}

// CallExpression represents a function call
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out string
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out += ce.Function.String() + "(" + strings.Join(args, ", ") + ")"
	return out
}

func (ce *CallExpression) expressionNode() {}

// BreakStatement represents a break statement
type BreakStatement struct {
	Token token.Token
}

func (bs *BreakStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BreakStatement) String() string {
	return "break;"
}

func (bs *BreakStatement) statementNode() {}

// ContinueStatement represents a continue statement
type ContinueStatement struct {
	Token token.Token
}

func (cs *ContinueStatement) TokenLiteral() string {
	return cs.Token.Literal
}

func (cs *ContinueStatement) String() string {
	return "continue;"
}

func (cs *ContinueStatement) statementNode() {}
