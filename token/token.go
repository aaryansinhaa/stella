package token

// TokenType is a string that represents the type of token
type TokenType string

// Token represents a lexical token with its type and literal value
type Token struct {
	Type    TokenType
	Literal string
}

// Token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	// Comparison operators
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	SET      = "SET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords maps string keywords to their corresponding TokenType
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"set":    SET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the TokenType for a given identifier
// tok is the token type if the identifier is a keyword
// otherwise it returns IDENT
// ok is true if the identifier is a keyword
func LookupIdent(ident string) TokenType {

	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
