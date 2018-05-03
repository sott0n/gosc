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
	FileName string `short:"f" long:"file" description:"interpret selected scheme source file"`
	DumpAST  bool   `short:"a" long:"ast" default:"false" description:"whether leafnodes are plotte"`
}

func main() {
	options := new(Options)
	if _, err := flags.Parse(options); err != nil {
		return
	}
	if len(options.FileName) > 0 {
		executeSourceCode(options)
	} else {
		repl(options)
	}
}

func executeSourceCode(options *Options) {
	buffer, err := ioutil.ReadFile(options.FileName)
	if err != nil {
		log.Fatal(err)
	}
	interpreter := scheme.NewInterpreter(string(buffer))
	interpreter.Eval(options.DumpAST)
}

func repl(options *Options) {
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
				interpreter := scheme.NewInterpreter(expression)
				interpreter.Eval(options.DumpAST)
				break
			} else if indentLevel < 0 {
				log.Println("Error: extra close parentheses.")
				expression = ""
				indentLevel = 0
			}
		}
	}
}

func callRepl(indentLevel int) string {
	if indentLevel == 0 {
		return "gosc>> "
	} else if indentLevel > 0 {
		return fmt.Sprintf("... %s", strings.Repeat("  ", indentLevel))
	} else {
		panic("Negative indent level.")
	}
}
