package repl

import (
	"bufio"
	"fmt"
	"go-int/src/lexer"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {

		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line:= scanner.Text()

		l:= lexer.New(line)
	}

}
