package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// the split function on bufio.Scanner
	// defaults to ScanLines -> till \n

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		l := lexer.New(line)

		// initialization; conditional; update
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}