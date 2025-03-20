package token

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Token types
const (
	ILLEGAL = "ILLEGAL" // unknown token
	EOF     = "EOF"     // end of file

	// Identifiers + literals
	IDENT  = "IDENT"  // add, x, y, ...
	INT    = "INT"    // 1, 2, 3, ...
	STRING = "STRING" // "foo", "bar", ...

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	LTE      = "<="
	GTE      = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	CONST    = "CONST"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	// Types
	TYPE_INT    = "INT_TYPE"
	TYPE_FLOAT  = "FLOAT_TYPE"
	TYPE_BOOL   = "BOOL_TYPE"
	TYPE_STRING = "STRING_TYPE"
	TYPE_CHAR   = "CHAR_TYPE"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"var":      VAR,
	"const":    CONST,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"true":     TRUE,
	"false":    FALSE,

	// Types
	"int":    TYPE_INT,
	"float":  TYPE_FLOAT,
	"bool":   TYPE_BOOL,
	"string": TYPE_STRING,
	"char":   TYPE_CHAR,
}

// LookupIdent checks if the given identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
