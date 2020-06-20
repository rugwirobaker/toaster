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

		// Misc characters
		{s: `*`, kind: ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, kind: IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, kind: IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, kind: FROM, lit: "FROM"},
		{s: `SELECT`, kind: SELECT, lit: "SELECT"},
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
