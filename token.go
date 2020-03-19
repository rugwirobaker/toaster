package toaster

import "strings"

// Kind ...
type Kind string

// Token ...
type Token struct {
	Literal string
	Kind    Kind
}

// tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//keywords
	CREATE = "CREATE"
	TABLE  = "TABLE"
	AS     = "AS"
	INSERT = "INSERT"
	SELECT = "SELECT"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	WHERE  = "WHERE"
	FROM   = "FROM"
	INTO   = "INTO"
	SET    = "SET"
	VALUES = "VALUES"
	BOOL   = "BOOL"
	TEXT   = "TEXT"
	INT    = "INT"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	// Identifiers
	IDENT = "IDENT"

	//datatypes
	NUMERIC = "NUMERIC"
	STRING  = "STRING"

	// Operators
	ASSIGN = "="

	// Delimiters
	ASTERIX   = "*"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
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

// LookupIdent ...
func LookupIdent(ident string) Kind {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
