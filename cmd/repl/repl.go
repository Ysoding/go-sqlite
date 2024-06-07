package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Ysoding/go-sqlite/internal/lexer"
	"github.com/Ysoding/go-sqlite/internal/parser"
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

		if strings.HasPrefix(line, ".") {
			doCommnd(line)
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)
		p.Parse()

	}
}

func printHelp() {
	fmt.Println("fake news!!!")
}

func doCommnd(cmd string) {
	switch cmd {
	case ".exit":
		os.Exit(0)
	case ".help":
		printHelp()
	}
}
