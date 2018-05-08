// This file is for statements by syntax form, such as set!

package scheme

// Set is a type for updating value.
type Set struct {
	ObjectBase
	variable Object
	value    Object
}

// NewSet is definition of creating new set pair.
func NewSet(parent Object) *Set {
	return &Set{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is for evaluating set object.
func (s *Set) Eval() Object {
	variable := s.variable.Eval()
	if !variable.isVariable() {
		variable = variable.Bounder()
		if variable == nil {
			compileError("syntax-error: malformed set!")
		}
	}
	value := s.value.Eval()
	s.ObjectBase.updateBinding(variable.(*Variable).identifier, value)
	return NewSymbol("#<undef>")
}

// If is for if statement object.
type If struct {
	ObjectBase
	condition Object
	trueBody  Object
	falseBody Object
}

// NewIf is definition of creating if statement.
func NewIf(parent Object) *If {
	return &If{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is IF of If statement's eval.
func (i *If) Eval() Object {
	result := i.condition.Eval()
	if result.isBoolean() && result.(*Boolean).value {
		return i.trueBody.Eval()
	}
	return i.falseBody.Eval()
}
