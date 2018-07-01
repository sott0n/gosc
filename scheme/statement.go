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
	return undef
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

// Cond is for cond statement object.
type Cond struct {
	ObjectBase
	cases    []Object
	elseBody Object
}

// NewCond is a new object for cond syntax.
func NewCond(parent Object) *Cond {
	return &Cond{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is evaluation for cond syntax.
func (c *Cond) Eval() Object {
	for _, caseBody := range c.cases {
		elements := caseBody.(*Pair).Elements()
		lastResult := elements[0].Eval()

		if lastResult.isBoolean() && lastResult.(*Boolean).value == false {
			continue
		}

		for _, element := range elements {
			lastResult = element.Eval()
		}
		return lastResult
	}

	if c.elseBody == nil {
		return undef
	}

	elements := c.elseBody.(*Pair).Elements()
	lastResult := Object(undef)
	for _, element := range elements {
		lastResult = element.Eval()
	}
	return lastResult
}

// And is a struct for and statement.
type And struct {
	ObjectBase
	body Object
}

// NewAnd is a new object for and syntax.
func NewAnd(parent Object) *And {
	return &And{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is evaluation for and syntax.
func (a *And) Eval() Object {
	lastResult := Object(NewBoolean(true))
	for _, object := range a.body.(*Pair).Elements() {
		lastResult = object.Eval()
		if lastResult.isBoolean() && lastResult.(*Boolean).value == false {
			return NewBoolean(false)
		}
	}
	return lastResult
}

// Or is a struct for or statement.
type Or struct {
	ObjectBase
	body Object
}

// NewOr is a new object for or syntax.
func NewOr(parent Object) *Or {
	return &Or{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is evaluation for or syntax.
func (o *Or) Eval() Object {
	lastResult := Object(NewBoolean(false))
	for _, object := range o.body.(*Pair).Elements() {
		lastResult = object.Eval()
		if !lastResult.isBoolean() || lastResult.(*Boolean).value != false {
			return lastResult
		}
	}
	return lastResult
}

// Begin is a struct for begin statement.
type Begin struct {
	ObjectBase
	body Object
}

// NewBegin is a new object for begin syntax.
func NewBegin(parent Object) *Begin {
	return &Begin{ObjectBase: ObjectBase{parent: parent}}
}

// Eval is evaluation for begin syntax.
func (b *Begin) Eval() Object {
	lastResult := Object(undef)
	for _, object := range b.body.(*Pair).Elements() {
		lastResult = object.Eval()
	}
	return lastResult
}

// Do is a struct for do statement.
type Do struct {
	ObjectBase
	iterators    []*Iterator
	testBody     Object
	continueBody Object
	localBinding Binding
}

// NewDo is a new object for do syntax.
func NewDo(parent Object) *Do {
	return &Do{ObjectBase: ObjectBase{parent: parent}, localBinding: Binding{}}
}

func (d *Do) binding() Binding {
	return d.localBinding
}

func (d *Do) scopedBinding() Binding {
	scopedBinding := make(Binding)
	for identifier, object := range d.localBinding {
		scopedBinding[identifier] = object
	}

	parent := d.Parent()
	for parent != nil {
		for identifier, object := range parent.binding() {
			if scopedBinding[identifier] == nil {
				scopedBinding[identifier] = object
			}
		}
		parent = parent.Parent()
	}
	return scopedBinding
}

// Eval is evaluation for do syntax.
func (d *Do) Eval() Object {
	// bind iterators
	for _, iterator := range d.iterators {
		if iterator.variable.isVariable() {
			d.localBinding[iterator.variable.(*Variable).identifier] = iterator.value.Eval()
		}
	}
	// eval test ->
	//   true: eval testBody and returns its result
	//   false: eval continueBody, eval iterator's update
	testElements := d.testBody.(*Pair).Elements()
	continueElements := d.continueBody.(*Pair).Elements()
	for {
		testResult := testElements[0].Eval()
		if !testResult.isBoolean() || testResult.(*Boolean).value == true {
			for _, element := range testElements[1:] {
				testResult = element.Eval()
			}
			return testResult
		} else {
			// eval continueBody
			for _, element := range continueElements {
				element.Eval()
			}

			// update iterators
			for _, iterator := range d.iterators {
				if iterator.variable.isVariable() {
					d.localBinding[iterator.variable.(*Variable).identifier] = iterator.update.Eval()
				}
			}
		}
	}
	return undef
}

// Iterator is a struct for iterator statement.
type Iterator struct {
	ObjectBase
	variable Object
	value    Object
	update   Object
}

// NewIterator is a new object for iterator syntax.
func NewIterator(parent Object) *Iterator {
	return &Iterator{ObjectBase: ObjectBase{parent: parent}}
}
