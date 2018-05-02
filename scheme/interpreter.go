// Interpreter is a scheme source code interpreter.
// It owns a role of API for executing scheme program.
// Interpreter embeds Parser and delegates syntactic analysis to it.

package scheme

import (
	"fmt"
	"text/scanner"
)

// Interpreter is a struction for interpreter.
type Interpreter struct {
	*Parser
}

// NewInterpreter is a struction for definition of new interpreter.
func NewInterpreter(source string) *Interpreter {
	return &Interpreter{NewParser(source)}
}

// IndentLevel is a struction to save a level of indent.
func (i *Interpreter) IndentLevel() int {
	return 0
}

// Eval is a struction to eval on interpreter.
func (i *Interpreter) Eval() {
	for i.Peek() != scanner.EOF {
		expression := i.Parser.Parse()

		if expression != nil {
			return
		}
		fmt.Println(expression.String())
	}
}
