// Scheme's identifier is classfied to a symbol or a variable.
// And this type own a role to express a variable.
// Variable itself does not have a value for identifier,
// interpreter searched it from its code block scope by Variable's identifier.

package scheme

// Variable is a struction for scheme variable object.
type Variable struct {
	ObjectBase
	identifier string
}

// NewVariable is a function for scheme variable object.
func NewVariable(identifier string, parent Object) *Variable {
	return &Variable{
		ObjectBase: ObjectBase{parent: parent},
		identifier: identifier,
	}
}

// Eval is variable's eval IF.
func (v *Variable) Eval() Object {
	object := v.boundedObject(v.identifier)
	if object == nil {
		runtimeError("Unbound variable: %s", v.identifier)
	}
	return object
}

func (v *Variable) String() string {
	return v.Eval().String()
}

func (v *Variable) isVariable() bool {
	return true
}
