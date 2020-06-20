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
		fmt.Printf(PROMPT)
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
			fmt.Printf("%+v\n", tok)
		}
	}
}
