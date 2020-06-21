package toaster

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var cases = []struct {
		input string
		kind  Kind
		lit   string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{input: ``, kind: EOF},
		{input: `#`, kind: ILLEGAL, lit: `#`},
		{input: ` `, kind: WS, lit: " "},
		{input: "\t", kind: WS, lit: "\t"},
		{input: "\n", kind: WS, lit: "\n"},
		{input: `,`, kind: COMMA, lit: ","},
		{input: `)`, kind: RPAREN, lit: ")"},
		{input: `(`, kind: LPAREN, lit: "("},
		{input: `;`, kind: SEMICOLON, lit: ";"},

		// Misc characters
		{input: `*`, kind: ASTERISK, lit: "*"},

		// Identifiers
		{input: `foo`, kind: IDENT, lit: `foo`},
		{input: `Zx12_3U_-`, kind: IDENT, lit: `Zx12_3U_`},
		{input: `users`, kind: IDENT, lit: "users"},
		{input: `(`, kind: LPAREN, lit: "("},
		{input: `id`, kind: IDENT, lit: "id"},

		// keywords
		{input: `bool`, kind: BOOL, lit: "bool"},
		{input: `CREATE`, kind: CREATE, lit: "create"},
		{input: `DELETE`, kind: DELETE, lit: "delete"},
		{input: `FROM`, kind: FROM, lit: "from"},
		{input: `INSERT`, kind: INSERT, lit: "insert"},
		{input: `INTO`, kind: INTO, lit: "into"},
		{input: `SELECT`, kind: SELECT, lit: "select"},
		{input: `TABLE`, kind: TABLE, lit: "table"},
		{input: `TEXT`, kind: TEXT, lit: "text"},
		{input: `VALUES`, kind: VALUES, lit: "values"},
		{input: `WHERE`, kind: WHERE, lit: "where"},
		{input: `INT`, kind: INT, lit: "int"},

		{input: `FALSE`, kind: FALSE, lit: "false"},
		{input: `TRUE`, kind: TRUE, lit: "true"},

		// // values
		// {input: `\'this is a string\'`, kind: STRING, lit: "this is a string"},
		{input: `1234`, kind: NUMBER, lit: "1234"},
		{input: `123.4`, kind: NUMBER, lit: "123.4"},
	}

	for i, tc := range cases {
		s := NewScanner(strings.NewReader(tc.input))
		tok := s.Scan()
		if tc.kind != tok.Kind {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tc.input, tc.kind, tok.Kind, tok.Literal)
		} else if tc.lit != tok.Literal {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tc.input, tc.lit, tok.Literal)
		}
	}
}
