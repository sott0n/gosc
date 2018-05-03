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
			expression += " "
			expression += currentLine

			interpreter := scheme.NewInterpreter(expression)
			indentLevel = interpreter.IndentLevel()
			if indentLevel == 0 {
				// Because the IndentLevel() method changes its reading position,
				// recreate intepreter to initialize the position.
				interpreter.Eval()
				break
			} else if indentLevel < 0 {
				log.Println("Error: extra close parentheses.")
				expression = ""
				indentLevel = 0
			}
		}
	}
}

func main() {
	repl()
}
