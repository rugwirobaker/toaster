package toaster

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// PROMPT output
const PROMPT = ">> "

// Start ...
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "\\q" {
			return
		}
		sc := NewScanner(strings.NewReader(line))
		for tok := sc.Scan(); tok.Kind != EOF; tok = sc.Scan() {
			fmt.Fprint(out, fmt.Sprintf("{Literal:%s Kind:%s}\n", tok.Literal, tok.String()))
		}
	}
}
