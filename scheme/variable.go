// Scheme's identifier is classfied to a symbol or a variable.
// And this type own a role to express a variable.
// Variable itself does not have a value for identifier,
// interpreter searched it from its code block scope by Variable's identifier.

package scheme

// Variable is a struction for scheme variable object.
type Variable struct {
	ObjectBase
	identifier  string
	environment *Environment
}

// NewVariable is a function for scheme variable object.
func NewVariable(identifier string, environment *Environment) *Variable {
	return &Variable{
		identifier:  identifier,
		environment: environment,
	}
}

// Eval is variable's eval IF.
func (v *Variable) Eval() Object {
	return v
}

func (v *Variable) String() string {
	return v.environment.binding[v.identifier].String()
}
