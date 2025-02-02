package token

// Token struct to represent a token with its type and literal.
type Token struct {
	Type, Literal string
}

// Constants representing the types of tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER         = "IDENTIFIER"
	FLOAT              = "FLOAT"
	INT                = "INT"
	STRING             = "STRING"
	BOOL               = "BOOL"
	UnterminatedString = "unterminated string"

	// Operators token
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"
	MODULO   = "MODULO"
	POWER    = "POWER"

	// Bitwise operators
	BITAND = "BITAND"
	BITOR  = "BITOR"
	BITNOT = "BITNOT"

	// Logical operators
	AND = "AND"
	OR  = "OR"

	// Comparisons token
	LESSTHAN     = "LESSTHAN"
	GREATERTHAN  = "GREATERTHAN"
	LESSEQUAL    = "LESSEQUAL"
	GREATEREQUAL = "GREATEREQUAL"
	EQUAL        = "EQUAL"
	NOT          = "NOT"
	NOTEQUAL     = "NOTEQUAL"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LEFTPARENTHESIS  = "LEFTPARENTHESIS"
	RIGHTPARENTHESIS = "RIGHTPARENTHESIS"
	LEFTBRACE        = "LEFTBRACE"
	RIGHTBRACE       = "RIGHTBRACE"
	LEFTBRACKET      = "LEFTBRACKET"
	RIGHTBRACKET     = "RIGHTBRACKET"

	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	CONST    = "CONST"
	TYPE     = "TYPE"
)

// Map keywords in the language to their respective token types.
var keywords = map[string]string{
	"func":   FUNCTION,
	"const":  CONST,
	"int":    TYPE,
	"string": TYPE,
	"bool":   TYPE,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Check if a given identifier is a keyword and return the corresponding type. Returns the IDENTIFIER type if not a keyword.
func LookupIdentifier(identifier string) string {
	tok, ok := keywords[identifier]
	if ok {
		return tok
	}
	return IDENTIFIER
}
