package toaster

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	input := `
		CREATE TABLE users (id INT, name TEXT, active BOOL);
		INSERT INTO users VALUES (1, 'Phil', TRUE);
		INSERT INTO users (id, name, active) VALUES (1, 'Phil', FALSE);
		SELECT * FROM users WHERE id=1;
		UPDATE users SET name='Jones' WHERE id=1;
		DELETE FROM users WHERE id=1;
	`

	// SELECT id, name FROM users where id=1;

	tests := []struct {
		expectedKind    Kind
		expectedLiteral string
	}{
		{CREATE, "create"},
		{TABLE, "table"},
		{IDENT, "users"},
		{LPAREN, "("},
		{IDENT, "id"},
		{INT, "int"},
		{COMMA, ","},
		{IDENT, "name"},
		{TEXT, "text"},
		{COMMA, ","},
		{IDENT, "active"},
		{BOOL, "bool"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{INSERT, "insert"},
		{INTO, "into"},
		{IDENT, "users"},
		{VALUES, "values"},
		{LPAREN, "("},
		{NUMERIC, "1"},
		{COMMA, ","},
		{STRING, "Phil"},
		{COMMA, ","},
		{TRUE, "true"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{INSERT, "insert"},
		{INTO, "into"},
		{IDENT, "users"},
		{LPAREN, "("},
		{IDENT, "id"},
		{COMMA, ","},
		{IDENT, "name"},
		{COMMA, ","},
		{IDENT, "active"},
		{RPAREN, ")"},
		{VALUES, "values"},
		{LPAREN, "("},
		{NUMERIC, "1"},
		{COMMA, ","},
		{STRING, "Phil"},
		{COMMA, ","},
		{FALSE, "false"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		{SELECT, "select"},
		{ASTERIX, "*"},
		{FROM, "from"},
		{IDENT, "users"},
		{WHERE, "where"},
		{IDENT, "id"},
		{ASSIGN, "="},
		{NUMERIC, "1"},
		{SEMICOLON, ";"},

		{UPDATE, "update"},
		{IDENT, "users"},
		{SET, "set"},
		{IDENT, "name"},
		{ASSIGN, "="},
		{STRING, "Jones"},
		{WHERE, "where"},
		{IDENT, "id"},
		{ASSIGN, "="},
		{NUMERIC, "1"},
		{SEMICOLON, ";"},

		{DELETE, "delete"},
		{FROM, "from"},
		{IDENT, "users"},
		{WHERE, "where"},
		{IDENT, "id"},
		{ASSIGN, "="},
		{NUMERIC, "1"},
		{SEMICOLON, ";"},
	}

	scanner := NewScanner(strings.NewReader(input))

	for i, tc := range tests {
		tok := scanner.Scan()

		if tok.Kind != tc.expectedKind {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tc.expectedKind, tok.Kind)
		}
		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tc.expectedLiteral, tok.Literal)
		}
	}
}
