package tokens

type TokenType string

type Token struct {
	Type TokenType

	Literal string
}

const (
	ILLEGAL = "Illegal"
	EOF     = "EOF"
	ASSIGN  = "="
	PLUS    = "+"
	IDENT   = "IDENT"
	//identifier
	INT       = "INT"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "}"
	RBRACE    = "{"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	ASTERISK  = "*"
	BANG      = "!"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT

}
