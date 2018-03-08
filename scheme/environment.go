// Environment has variable bindings.
// Interpreter has one Environment global variable for top-level environment.
// And each let block and procedure has Environment to hold its scope's variable binding.

package scheme

// Environment is a struction for environment.
type Environment struct {
	ObjectBase
	parent  *Environment
	binding *Binding
}

// Binding is a struction for binding.
type Binding map[string]*Procedure

func newEnvironment() *Environment {
	return &Environment{}
}
