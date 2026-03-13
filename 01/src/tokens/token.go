package tokens

type TokenType string

type Token struct {
	Type TokenType

	Literal string
}

const (
	ILLEGAL   = "Illegal"
	EOF       = "EOF"
	ASSIGN    = "="
	PLUS      = "+"
	IDENT     = "IDENT"
	//identifier
	INT       = "INT"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "}"
	RBRACE    = "{"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)