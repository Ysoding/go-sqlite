package token

// https://github.com/sqlite/sqlite/blob/master/src/tokenize.c
const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTARISK = "*"
	SLASH    = "/"

	LT  = "<"
	GT  = ">"
	EQ  = "=="
	NEQ = "!="

	LPAREN = "("
	RPAREN = ")"

	COMMA = ","

	SELECT = "SELECT"
	DELETE = "DELETE"
	INSERT = "INSERT"
	FROM   = "FROM"

	IDENT = "IDENT"
	INT   = "INT"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"select": SELECT,
	"delete": DELETE,
	"insert": INSERT,
	"from":   FROM,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
