package toaster

import (
	"strings"
)

// Scanner ...
type Scanner struct {
	input    string
	position int // current position in input (points to current char)
	next     int // current reading position in input (after current char)
	ch       byte
}

// New scanner
func New(input string) *Scanner {
	scanner := &Scanner{input: input}
	scanner.readChar()
	return scanner
}

// NextToken ...
func (sc *Scanner) NextToken() Token {
	var tok Token

	sc.skipWhitespace()

	switch sc.ch {
	case '*':
		tok = newToken(ASTERIX, sc.ch)
	case '=':
		tok = newToken(ASSIGN, sc.ch)
	case ';':
		tok = newToken(SEMICOLON, sc.ch)
	case '(':
		tok = newToken(LPAREN, sc.ch)
	case ')':
		tok = newToken(RPAREN, sc.ch)
	case ',':
		tok = newToken(COMMA, sc.ch)
	case 0:
		tok.Literal = ""
		tok.Kind = EOF
	case '\'':
		tok.Kind = STRING
		tok.Literal = sc.readString()
	default:
		if isLetter(sc.ch) {
			tok.Literal = sc.readIdentifier()
			tok.Kind = LookupIdent(tok.Literal)
			if tok.Kind != IDENT {
				tok.Literal = strings.ToLower(tok.Literal)
			}
			return tok
		}
		if isDigit(sc.ch) {
			tok.Kind = NUMERIC
			tok.Literal = sc.readNumber()
			return tok
		}
		tok = newToken(ILLEGAL, sc.ch)

	}

	sc.readChar()
	return tok
}

func newToken(tokenKind Kind, ch byte) Token {
	return Token{Kind: tokenKind, Literal: string(ch)}
}

func (sc *Scanner) readChar() {
	// check if we have reached the end of input
	if sc.next >= len(sc.input) {
		sc.ch = 0
	} else {
		sc.ch = sc.input[sc.next]
	}
	sc.position = sc.next
	sc.next++
}

func (sc *Scanner) readIdentifier() string {
	position := sc.position
	for isLetter(sc.ch) {
		sc.readChar()
	}
	return sc.input[position:sc.position]
}

func (sc *Scanner) skipWhitespace() {
	for sc.ch == ' ' || sc.ch == '\t' || sc.ch == '\n' || sc.ch == '\r' {
		sc.readChar()
	}
}

func (sc *Scanner) readNumber() string {
	position := sc.position
	for isDigit(sc.ch) {
		sc.readChar()
	}
	return sc.input[position:sc.position]
}

func (sc *Scanner) readString() string {
	position := sc.position + 1
	for {
		sc.readChar()
		if sc.ch == '\'' || sc.ch == 0 {
			break
		}
	}
	return sc.input[position:sc.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
