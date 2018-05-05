package main

import (
	"fmt"
	"gosc/scheme"
	"io/ioutil"
	"log"
	"strings"

	"github.com/GeertJohan/go.linenoise"
	"github.com/jessevdk/go-flags"
)

// Options is definition gosc's selection option.
type Options struct {
	FileName   string   `short:"f" long:"file" description:"interpret selected scheme source file"`
	Expression []string `short:"e" long:"expression" description:"execute given expression"`
	DumpAST    bool     `short:"a" long:"ast" description:"whether leaf nodes are plotted"`
}

func main() {
	options := new(Options)
	if _, err := flags.Parse(options); err != nil {
		return
	}
	if len(options.FileName) > 0 {
		executeSourceCode(options)
	} else if len(options.Expression) > 0 {
		executeExpression(strings.Join(options.Expression, " "), options.DumpAST)
	} else {
		repl(options)
	}
}

func executeSourceCode(options *Options) {
	buffer, err := ioutil.ReadFile(options.FileName)
	if err != nil {
		log.Fatal(err)
	}
	executeExpression(string(buffer), options.DumpAST)
}

func executeExpression(expression string, dumpAST bool) {
	interpreter := scheme.NewInterpreter(expression)
	interpreter.PrintResult(dumpAST)
}

func repl(options *Options) {
	fmt.Println(">>> REPL of gosc is running...")
	mainInterpreter := scheme.NewInterpreter("")

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
			linenoise.AddHistory(currentLine)
			expression += " "
			expression += currentLine

			interpreter := scheme.NewInterpreter(expression)
			indentLevel = interpreter.IndentLevel()
			if indentLevel == 0 {
				mainInterpreter.ReloadSourceCode(expression)
				mainInterpreter.PrintResult(options.DumpAST)
				break
			} else if indentLevel < 0 {
				fmt.Println("*** ERROR: extra close parentheses")
				expression = ""
				indentLevel = 0
			}
		}
	}
}

func callRepl(indentLevel int) string {
	if indentLevel == 0 {
		return ">>> "
	} else if indentLevel > 0 {
		return fmt.Sprintf("... %s", strings.Repeat("  ", indentLevel))
	} else {
		panic("Negative indent level")
	}
}
