package toaster

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s    string
		kind Kind
		lit  string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, kind: EOF},
		{s: `#`, kind: ILLEGAL, lit: `#`},
		{s: ` `, kind: WS, lit: " "},
		{s: "\t", kind: WS, lit: "\t"},
		{s: "\n", kind: WS, lit: "\n"},
		{s: `;`, kind: SEMICOLON, lit: ";"},

		// Misc characters
		{s: `*`, kind: ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, kind: IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, kind: IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `CREATE`, kind: CREATE, lit: "CREATE"},
		{s: `TABLE`, kind: TABLE, lit: "TABLE"},
		{s: `INSERT`, kind: INSERT, lit: "INSERT"},
		{s: `SELECT`, kind: SELECT, lit: "SELECT"},
		{s: `UPDATE`, kind: UPDATE, lit: "UPDATE"},
		{s: `DELETE`, kind: DELETE, lit: "DELETE"},
		{s: `FROM`, kind: FROM, lit: "FROM"},
		{s: `WHERE`, kind: WHERE, lit: "WHERE"},
		{s: `INTO`, kind: INTO, lit: "INTO"},
		{s: `SET`, kind: SET, lit: "SET"},
		{s: `AS`, kind: AS, lit: "AS"},
		{s: `VALUES`, kind: VALUES, lit: "VALUES"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok := s.Scan()
		if tt.kind != tok.Kind {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.kind, tok.Kind, tok.Literal)
		} else if tt.lit != tok.Literal {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, tok.Literal)
		}
	}
}
