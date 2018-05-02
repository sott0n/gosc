// Interpreter is a scheme source code interpreter.
// It owns a role of API for executing scheme program.
// Interpreter embeds Parser and delegates syntactic analysis to it.

package scheme

import (
	"fmt"
	"strings"
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

// DumpAST is a defining of dumping abstrct tree.
func (i *Interpreter) DumpAST(object Object) {
	i.dumpASTWithIndent(object, 0)
}

// dumpASTWithIndent is a dumping AST function with specified indent.
func (i *Interpreter) dumpASTWithIndent(object Object, indentLevel int) {
	if object == nil {
		return
	}
	switch object.(type) {
	case *Application:
		i.printWithIndent("Application", indentLevel)
		i.dumpASTWithIndent(object.(*Application).procedureVariable, indentLevel+1)
		i.dumpASTWithIndent(object.(*Application).arguments, indentLevel+1)
	case *Pair:
		pair := object.(*Pair)
		if pair.Car == nil && pair.Cdr == nil {
			return
		}
		i.printWithIndent("Pair", indentLevel)
		i.dumpASTWithIndent(pair.Car, indentLevel+1)
		i.dumpASTWithIndent(pair.Cdr, indentLevel+1)
	case *Number:
		i.printWithIndent(fmt.Sprintf("Number(%d)", object.(*Number).value), indentLevel)
	case *Variable:
		i.printWithIndent(fmt.Sprintf("Variable(%s)", object.(*Variable).identifier), indentLevel)
	}
}

func (i *Interpreter) printWithIndent(text string, indentLevel int) {
	fmt.Printf("%s%s\n", strings.Repeat(" ", indentLevel), text)
}
