package parser

import (
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/ast"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"testing"
)

func TestParseIntNode(t *testing.T) {
	input := "42"
	lexer := lexer.New(input, lexer.DefaultConfig) // Ensure this is your Lexer
	parser := New(lexer)                           // Update to match `New` for `Parser`

	node := parser.parseIntLiteral()

	if node == nil {
		t.Fatalf("parseIntLiteral returned nil")
	}

	intNode, ok := node.(*ast.IntNode)
	if !ok {
		t.Fatalf("expected *ast.IntNode, got %T", node)
	}

	if intNode.Value != 42 {
		t.Errorf("expected value 42, got %d", intNode.Value)
	}
}
