package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Ysoding/go-sqlite/parser"
)

const PROMPT = "db > "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(line)
		p.Parse(out)
	}
}
