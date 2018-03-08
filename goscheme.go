package main

import (
	"fmt"
	"go-scheme/scheme"
	"log"
	"strings"

	"github.com/GeertJohan/go.linenoise"
)

func callRepl(indentLevel int) string {
	if indentLevel == 0 {
		return "gosc>> "
	} else if indentLevel > 0 {
		return fmt.Sprintf("... %s", strings.Repeat("  ", indentLevel))
	} else {
		panic("Negative indent level.")
	}
}

func repl() {
	for {
		indentLevel := 0
		expression := ""
		for {
			currentLine, err := linenoise.Line(callRepl(indentLevel))
			if err != nil {
				log.Fatal(err)
				return
			}
			if currentLine == "exit" {
				return
			}
			expression += currentLine

			interpreter := scheme.NewInterpreter(expression)
			if indentLevel = interpreter.IndentLevel(); indentLevel == 0 {
				interpreter.Eval()
				break
			}
		}
	}
}

func main() {
	repl()
}
