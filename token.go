package toaster

import (
	"strings"
)

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

	//boolean literals
	TRUE
	FALSE

	// Operators
	ASSIGN

	// Delimiters
	ASTERISK
	COMMA
	SEMICOLON
	LPAREN
	RPAREN

	keywordBeg
	//keywords
	AS
	BOOL
	CREATE
	DELETE
	FROM
	INSERT
	INT
	INTO
	SELECT
	SET
	TABLE
	TEXT
	UPDATE
	WHERE
	VALUES
	keywordEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	IDENT:     "IDENT",
	NUMBER:    "NUMBER",
	STRING:    "STRING",
	BADSTRING: "BADSTRING",
	BADESCAPE: "BADESCAPE",

	TRUE:  "true",
	FALSE: "false",

	ASSIGN: "=",

	ASTERISK:  "*",
	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",

	AS:     "as",
	BOOL:   "bool",
	CREATE: "create",
	DELETE: "delete",
	FROM:   "from",
	INSERT: "insert",
	INT:    "int",
	INTO:   "into",
	SELECT: "select",
	SET:    "set",
	TABLE:  "table",
	TEXT:   "text",
	UPDATE: "update",
	VALUES: "values",
	WHERE:  "where",
}

var keywords map[string]Kind

func init() {
	keywords = make(map[string]Kind)

	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}

	keywords["true"] = TRUE
	keywords["false"] = FALSE
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok.Kind >= 0 && tok.Kind < Kind(len(tokens)) {
		return tokens[tok.Kind]
	}
	return ""
}

// LookupIdent ...
func LookupIdent(ident string) Kind {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func newToken(kind Kind, lit string) Token {
	return Token{Kind: kind, Literal: lit}
}
