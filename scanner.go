package toaster

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

var eof = rune(0)

// Scanner ...
type Scanner struct {
	r *bufio.Reader
}

// NewScanner instantiates a new scanner instance with input
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter, or certain acceptable special characters, then consume
	// as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	switch ch {
	case '*':
		return newToken(ASTERISK, string(ch))
	case '=':
		return newToken(ASSIGN, string(ch))
	case ';':
		return newToken(SEMICOLON, string(ch))
	case '(':
		return newToken(LPAREN, string(ch))
	case ')':
		return newToken(RPAREN, string(ch))
	case ',':
		return newToken(COMMA, string(ch))
	case '\'':
		s.unread()
		return s.scanString()
	case eof:
		return newToken(EOF, "")
	default:
		if isLetter(ch) {
			s.unread()
			return s.scanIdent()
		}
		if isDigit(ch) {
			s.unread()
			return s.scanNumber()
		}
	}
	return newToken(ILLEGAL, string(ch))
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return newToken(WS, buf.String())
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	lit := buf.String()

	if kind := LookupIdent(lit); kind != IDENT {
		return newToken(kind, strings.ToLower(lit))
	}
	return newToken(IDENT, lit)
}

func (s *Scanner) scanNumber() (tok Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) && ch != '.' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	lit := buf.String()

	return newToken(NUMBER, lit)
}

func (s *Scanner) scanString() (tok Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	s.unread()

	lit, err := ScanString(s.r)
	if err == errBadString {
		return newToken(BADSTRING, lit)
	} else if err == errBadEscape {
		return newToken(BADESCAPE, lit)
	}
	return newToken(STRING, lit)
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()

	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// ScanString reads a quoted string from a rune reader.
func ScanString(r *bufio.Reader) (string, error) {
	ending, _, err := r.ReadRune()
	if err != nil {
		return "", errBadString
	}

	var buf bytes.Buffer
	for {
		ch, _, err := r.ReadRune()
		if ch == ending {
			return buf.String(), nil
		} else if err != nil || ch == '\n' {
			return buf.String(), errBadString
		} else if ch == '\\' {
			// If the next character is an escape then write the escaped char.
			// If it's not a valid escape then return an error.
			ch1, _, _ := r.ReadRune()
			if ch1 == 'n' {
				_, _ = buf.WriteRune('\n')
			} else if ch1 == '\\' {
				_, _ = buf.WriteRune('\\')
			} else if ch1 == '"' {
				_, _ = buf.WriteRune('"')
			} else if ch1 == '\'' {
				_, _ = buf.WriteRune('\'')
			} else {
				return string(ch) + string(ch1), errBadEscape
			}
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
}

var errBadString = errors.New("bad string")
var errBadEscape = errors.New("bad escape")
