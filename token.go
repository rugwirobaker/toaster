package toaster

import "strings"

// Kind ...
type Kind int

// Token ...
type Token struct {
	Literal string
	Kind    Kind
}

// tokens
const (
	ILLEGAL Kind = iota
	EOF
	WS
	COMMENT

	// Literals
	IDENT
	NUMBER    // 12345.67
	STRING    // "abc"
	BADSTRING // "abc
	BADESCAPE // \q

	// Operators
	ASSIGN

	// Delimiters
	ASTERISK
	COMMA
	SEMICOLON
	LPAREN
	RPAREN

	//keywords
	CREATE
	TABLE
	AS
	INSERT
	SELECT
	UPDATE
	DELETE
	WHERE
	FROM
	INTO
	SET
	VALUES
	BOOL
	TEXT
	INT
	TRUE
	FALSE
)

var keywords = map[string]Kind{
	"create": CREATE,
	"table":  TABLE,
	"insert": INSERT,
	"select": SELECT,
	"update": UPDATE,
	"delete": DELETE,
	"where":  WHERE,
	"from":   FROM,
	"into":   INTO,
	"set":    SET,
	"as":     AS,
	"values": VALUES,
	"text":   TEXT,
	"int":    INT,
	"bool":   BOOL,
	"true":   TRUE,
	"false":  FALSE,
}

func newToken(kind Kind, lit string) Token {
	return Token{Kind: kind, Literal: lit}
}

// LookupIdent ...
func LookupIdent(ident string) Kind {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
