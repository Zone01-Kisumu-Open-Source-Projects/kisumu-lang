package token

// Token struct to represent a token with its type and literal.
type Token struct {
	Type, Literal string
}

// Constants representing the types of tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	STRING     = "STRING"

	// Operators token
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	ASTERIK = "*"
	SLASH   = "/"

	// Comparisons token
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	EQUAL        = "=="
	NOT          = "!"
	NOT_EQUAL    = "!="

	COMMA      = ","
	SEMI_COLON = ";"
	COLON      = ":"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"
	LEFT_BRACKET      = "["
	RIGHT_BRACKET     = "]"

	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)