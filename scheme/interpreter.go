// Interpreter is a scheme source code interpreter.
// It owns a role of API for executing scheme program.
// Interpreter embeds Parser and delegates syntactic analysis to it.

package scheme

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/scanner"
)

// Interpreter is a struction for interpreter.
type Interpreter struct {
	ObjectBase
	*Parser
	topLevel Binding
}

// NewInterpreter is a struction for definition of new interpreter.
func NewInterpreter(source string) *Interpreter {
	i := &Interpreter{Parser: NewParser(source), topLevel: BuiltinProcedures()}
	i.loadBuiltinLibrary("builtin")
	return i
}

// ReloadSourceCode is to load new source code with current environment.
func (i *Interpreter) ReloadSourceCode(source string) {
	i.Parser = NewParser(source)
}

// PrintResult is a function to print result of Eval.
func (i *Interpreter) PrintResult(dumpAST bool) {
	results := i.EvalSource(dumpAST)
	if dumpAST {
		fmt.Printf("\n*** Result ***\n")
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

// EvalSource is a struction to eval on interpreter.
func (i *Interpreter) EvalSource(dumpAST bool) (results []string) {
	defer func() {
		if err := recover(); err != nil {
			results = append(results, fmt.Sprintf("*** ERROR: %s", err))
		}
	}()

	for i.Peek() != scanner.EOF {
		expression := i.Parser.Parse(i)
		if dumpAST {
			fmt.Printf("\n*** AST ***\n")
			i.DumpAST(expression, 0)
			fmt.Printf("\n*** Result ***\n")
		}

		if expression == nil {
			return
		}
		results = append(results, expression.Eval().String())
	}
	return
}

// DumpAST is a defining of dumping abstrct tree.
func (i *Interpreter) DumpAST(object Object, indentLevel int) {
	if object == nil {
		return
	}
	switch object.(type) {
	case *Application:
		i.printWithIndent("Application", indentLevel)
		i.DumpAST(object.(*Application).procedure, indentLevel+1)
		i.DumpAST(object.(*Application).arguments, indentLevel+1)
	case *Pair:
		pair := object.(*Pair)
		if pair.Car == nil && pair.Cdr == nil {
			return
		}
		i.printWithIndent("Pair", indentLevel)
		i.DumpAST(pair.Car, indentLevel+1)
		i.DumpAST(pair.Cdr, indentLevel+1)
	case *String:
		i.printWithIndent(fmt.Sprintf("String(%s)", object), indentLevel)
	case *Number:
		i.printWithIndent(fmt.Sprintf("Number(%s)", object), indentLevel)
	case *Boolean:
		i.printWithIndent(fmt.Sprintf("Boolean(%s)", object), indentLevel)
	case *Variable:
		i.printWithIndent(fmt.Sprintf("Variable(%s)", object.(*Variable).identifier), indentLevel)
	case *Definition:
		i.printWithIndent("Difinition", indentLevel)
		i.DumpAST(object.(*Definition).variable, indentLevel+1)
		i.DumpAST(object.(*Definition).value, indentLevel+1)
	case *Procedure:
		i.printWithIndent("Procedure", indentLevel)
		i.DumpAST(object.(*Procedure).arguments, indentLevel+1)
		i.DumpAST(object.(*Procedure).body, indentLevel+1)
	case *Set:
		i.printWithIndent("Set", indentLevel)
		i.DumpAST(object.(*Set).variable, indentLevel+1)
		i.DumpAST(object.(*Set).value, indentLevel+1)
	case *If:
		i.printWithIndent("If", indentLevel)
		i.DumpAST(object.(*If).condition, indentLevel+1)
		i.DumpAST(object.(*If).trueBody, indentLevel+1)
		i.DumpAST(object.(*If).falseBody, indentLevel+1)
	case *Cond:
		i.printWithIndent("Cond", indentLevel)
		for _, element := range object.(*Cond).cases {
			i.DumpAST(element, indentLevel+1)
		}
		if object.(*Cond).elseBody != nil {
			i.DumpAST(object.(*Cond).elseBody, indentLevel+1)
		}
	case *And:
		i.printWithIndent("And", indentLevel)
		i.DumpAST(object.(*And).body, indentLevel+1)
	case *Or:
		i.printWithIndent("Or", indentLevel)
		i.DumpAST(object.(*Or).body, indentLevel+1)
	case *Begin:
		i.printWithIndent("Begin", indentLevel)
		i.DumpAST(object.(*Begin).body, indentLevel+1)
	case *Do:
		i.printWithIndent("Do", indentLevel)
		for _, iterator := range object.(*Do).iterators {
			i.DumpAST(iterator, indentLevel+1)
		}
		i.DumpAST(object.(*Do).testBody, indentLevel+1)
		i.DumpAST(object.(*Do).continueBody, indentLevel+1)
	case *Iterator:
		i.printWithIndent("Iterator", indentLevel)
		i.DumpAST(object.(*Iterator).variable, indentLevel+1)
		i.DumpAST(object.(*Iterator).value, indentLevel+1)
		i.DumpAST(object.(*Iterator).update, indentLevel+1)
	}
}

func (i *Interpreter) bind(identifier string, object Object) {
	i.topLevel[identifier] = object
}

func (i *Interpreter) binding() Binding {
	return i.topLevel
}

func (i *Interpreter) scopedBinding() Binding {
	return i.topLevel
}

func (i *Interpreter) printWithIndent(text string, indentLevel int) {
	fmt.Printf("%s%s\n", strings.Repeat(" ", indentLevel), text)
}

func (i *Interpreter) loadBuiltinLibrary(name string) {
	originalParser := i.Parser
	buffer, err := ioutil.ReadFile(i.libraryPath(name))
	if err != nil {
		log.Fatal(err)
	}
	i.Parser = NewParser(string(buffer))
	i.EvalSource(false)
	i.Parser = originalParser
}

func (i *Interpreter) libraryPath(name string) string {
	return filepath.Join(
		os.Getenv("GOPATH"),
		"src",
		"gosc",
		"lib",
		name+".scm",
	)
}
