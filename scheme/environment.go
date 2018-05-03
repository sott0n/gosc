// Environment has variable bindings.
// Interpreter has one Environment global variable for top-level environment.
// And each let block and procedure has Environment to hold its scope's variable binding.

package scheme

import "log"

// Environment is a struction for environment.
type Environment struct {
	ObjectBase
	parent  *Environment
	binding Binding
}

// Binding is a struction for binding.
type Binding map[string]Object

// TopLevel is a setting environment.
var TopLevel = Environment{
	parent:  nil,
	binding: builtinProcedures,
}

func newEnvironment() *Environment {
	return &Environment{}
}

// Bind is to bind identifier and value in environment.
func (e *Environment) Bind(identifier string, value Object) {
	e.binding[identifier] = value
}

// Search procedure which is binded with given variable from environment,
// and invoke the procedure with given arguments.
func (e *Environment) invokeProcedure(variable, arguments Object) Object {
	if variable == nil {
		log.Fatal("Invoke procedure for <nil> variable.")
	}
	identifier := variable.(*Variable).identifier
	procedure := e.binding[identifier].(*Procedure)
	if procedure == nil {
		log.Printf("Unbound variable: %s\n", identifier)
	}
	return procedure.invoke(arguments)
}

func (e *Environment) scopedBinding() Binding {
	return Binding{}
}
