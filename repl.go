package toaster

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
			os.Exit(0)
		}
		sc := New(line)
		for tok := sc.NextToken(); tok.Kind != EOF; tok = sc.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
