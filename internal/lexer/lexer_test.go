package lexer_test

import (
	"testing"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

func TestLexerEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "Empty input",
			input: "",
			expected: []token.Token{
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Unterminated string",
			input: `"unclosed string`,
			expected: []token.Token{
				{Type: token.ILLEGAL, Literal: "unterminated string"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Unterminated multiline comment",
			input: "/* unclosed comment",
			expected: []token.Token{
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Numbers with decimal points",
			input: "42.5 3.14159 .5 5.",
			expected: []token.Token{
				{Type: token.FLOAT, Literal: "42.5"},
				{Type: token.FLOAT, Literal: "3.14159"},
				{Type: token.FLOAT, Literal: ".5"},
				{Type: token.FLOAT, Literal: "5."},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Complex nested expressions",
			input: "((x + y) * (z - 1)) / 2",
			expected: []token.Token{
				{Type: token.LEFTPARENTHESIS, Literal: "("},
				{Type: token.LEFTPARENTHESIS, Literal: "("},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENTIFIER, Literal: "y"},
				{Type: token.RIGHTPARENTHESIS, Literal: ")"},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.LEFTPARENTHESIS, Literal: "("},
				{Type: token.IDENTIFIER, Literal: "z"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.INT, Literal: "1"},
				{Type: token.RIGHTPARENTHESIS, Literal: ")"},
				{Type: token.RIGHTPARENTHESIS, Literal: ")"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.INT, Literal: "2"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "Mixed comments and code",
			input: `// Line comment
			x = 1; /* Multi
			line comment */ y = 2;
			// Another comment`,
			expected: []token.Token{
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "1"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.IDENTIFIER, Literal: "y"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "2"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Special characters and operators",
			input: "!= == < <= > >= && ||",
			expected: []token.Token{
				{Type: token.NOTEQUAL, Literal: "!="},
				{Type: token.EQUAL, Literal: "=="},
				{Type: token.LESSTHAN, Literal: "<"},
				{Type: token.LESSEQUAL, Literal: "<="},
				{Type: token.GREATERTHAN, Literal: ">"},
				{Type: token.GREATEREQUAL, Literal: ">="},
				{Type: token.AND, Literal: "&&"},
				{Type: token.OR, Literal: "||"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input, lexer.DefaultConfig)

			for i, expected := range tt.expected {
				got := l.NextToken()
				if got.Type != expected.Type || got.Literal != expected.Literal {
					t.Errorf("test %q - token[%d] wrong. expected=%+v, got=%+v",
						tt.name, i, expected, got)
				}
			}
		})
	}
}

func TestLexerConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		config   lexer.Config
		expected []token.Token
	}{
		{
			name:  "Disabled comments",
			input: "x = 1; // comment \n y = 2;",
			config: lexer.Config{
				AllowComments:    false,
				StringDelimiters: []byte{'"'},
			},
			expected: []token.Token{
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "1"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.IDENTIFIER, Literal: "comment"},
				{Type: token.IDENTIFIER, Literal: "y"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "2"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Custom string delimiters",
			input: "'single' \"double\"",
			config: lexer.Config{
				AllowComments:    true,
				StringDelimiters: []byte{'\'', '"'},
			},
			expected: []token.Token{
				{Type: token.STRING, Literal: "single"},
				{Type: token.STRING, Literal: "double"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input, tt.config)

			for i, expected := range tt.expected {
				got := l.NextToken()
				if got.Type != expected.Type || got.Literal != expected.Literal {
					t.Errorf("test %q - token[%d] wrong. expected=%+v, got=%+v",
						tt.name, i, expected, got)
				}
			}
		})
	}
}
