package parser

import (
	"testing"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/ast"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
)

func TestParseIntNode(t *testing.T) {
	input := "42"
	l := lexer.New(input, lexer.DefaultConfig)
	p := New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		t.Fatalf("parser has %d errors", len(p.Errors()))
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	intNode, ok := stmt.Expression.(*ast.IntNode)
	if !ok {
		t.Fatalf("exp not *ast.IntNode. got=%T", stmt.Expression)
	}

	if intNode.Value != 42 {
		t.Errorf("intNode.Value not %d. got=%d", 42, intNode.Value)
	}
}
