package toaster

import (
	"bytes"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{in: `CREATE`, out: `{Literal:create Kind:CREATE}`},
		{in: `TABLE`, out: `{Literal:table Kind:TABLE}`},
		{in: `users`, out: `{Literal:users Kind:IDENT}`},
	}

	for _, tc := range cases {
		var buf bytes.Buffer

		Start(strings.NewReader(tc.in), &buf)

		_, err := buf.Read([]byte(PROMPT))
		if err != nil {
			t.Errorf("unable to read prompt %v", err)
		}

		_, err = buf.Read([]byte(tc.out))
		if err != nil {
			t.Errorf("unable to read output %v", err)
		}
	}
}
