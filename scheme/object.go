// Object and ObjectBase is an abstruct class for all scheme expression.
// A return value of a method which returns scheme object is Object.
// And ObjectBase has Object's implementation of String().

package scheme

// Object is an abstruct class for scheme object.
type Object interface {
	Parent() Object
	Bounder() *Variable
	setParent(Object)
	setBounder(*Variable)
	Eval() Object
	String() string
	isNumber() bool
	isBoolean() bool
	isProcedure() bool
	isNull() bool
	isPair() bool
	isList() bool
	isSymbol() bool
	isString() bool
	isVariable() bool
	isApplication() bool
	bind(string, Object)
	updateBinding(string, Object)
	scopedBinding() Binding
	binding() Binding
	boundedObject(string) Object
	ancestor() Object
}

// Binding is an abstruct type for binding.
type Binding map[string]Object

// ObjectBase is an abstruct class for base scheme object.
type ObjectBase struct {
	parent  Object
	bounder *Variable // Variable.Eval() sets itself into this
}

// Eval is object's eval IF.
func (o *ObjectBase) Eval() Object {
	runtimeError("This type's String() is not implemented yet")
	return nil
}

func (o *ObjectBase) String() string {
	runtimeError("This object's String() is not implemented yet")
	return ""
}

func (o *ObjectBase) isNumber() bool {
	return false
}

func (o *ObjectBase) isBoolean() bool {
	return false
}

func (o *ObjectBase) isProcedure() bool {
	return false
}

func (o *ObjectBase) isNull() bool {
	return false
}

func (o *ObjectBase) isPair() bool {
	return false
}

func (o *ObjectBase) isList() bool {
	return false
}

func (o *ObjectBase) isVariable() bool {
	return false
}

func (o *ObjectBase) isSymbol() bool {
	return false
}

func (o *ObjectBase) isString() bool {
	return false
}

func (o *ObjectBase) isApplication() bool {
	return false
}

func (o *ObjectBase) binding() Binding {
	return Binding{}
}

// Bounder is IF that returns object's bounder.
func (o *ObjectBase) Bounder() *Variable {
	return o.bounder
}

// Parent is an abstruct function for accessing parent.
func (o *ObjectBase) Parent() Object {
	return o.parent
}

func (o *ObjectBase) setParent(parent Object) {
	o.parent = parent
}

func (o *ObjectBase) setBounder(bounder *Variable) {
	o.bounder = bounder
}

func (o *ObjectBase) scopedBinding() (scopedBinding Binding) {
	scopedBinding = make(Binding)
	parent := o.Parent()

	for parent != nil {
		for identifier, object := range parent.binding() {
			if scopedBinding[identifier] == nil {
				scopedBinding[identifier] = object
			}
		}
		parent = parent.Parent()
	}
	return
}

// This is for define syntax.
// Define variable in the top level.
func (o *ObjectBase) bind(identifier string, object Object) {
	if o.parent == nil {
		runtimeError("Bind called for object whose parent is nil")
	}
	o.ancestor().bind(identifier, object)
}

// This is for set! syntax.
// Update the variable's value when it is defined.
func (o *ObjectBase) updateBinding(identifier string, object Object) {
	target := o.Parent()
	for {
		if target == nil {
			break
		}
		if target.binding()[identifier] != nil {
			target.binding()[identifier] = object
			return
		}
		target = target.Parent()
	}
	runtimeError("symbol not defined")
}

func (o *ObjectBase) boundedObject(identifier string) Object {
	return o.scopedBinding()[identifier]
}

func (o *ObjectBase) ancestor() Object {
	ancestor := o.Parent()
	for {
		if ancestor.Parent() == nil {
			break
		} else {
			ancestor = ancestor.Parent()
		}
	}
	return ancestor
}
