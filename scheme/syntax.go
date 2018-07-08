// This file is for statements by syntax form, such as set!

package scheme

import "fmt"

var (
	builtinSyntaxes = Binding{
		"set!":   NewSyntax(setSyntax),
		"if":     NewSyntax(ifSyntax),
		"and":    NewSyntax(andSyntax),
		"or":     NewSyntax(orSyntax),
		"begin":  NewSyntax(beginSyntax),
		"define": NewSyntax(defineSyntax),
	}
)

// Syntax is a type for updating value.
type Syntax struct {
	ObjectBase
	function func(*Syntax, Object) Object
}

// NewSyntax is definition of creating new set pair.
func NewSyntax(function func(*Syntax, Object) Object) *Syntax {
	return &Syntax{ObjectBase: ObjectBase{parent: nil}, function: function}
}

// Invoke is for evaluating set object.
func (s *Syntax) Invoke(arguments Object) Object {
	return s.function(s, arguments)
}

func (s *Syntax) String() string {
	return fmt.Sprintf("#<syntax %s>", s.Bounder())
}

func (s *Syntax) isSyntax() bool {
	return true
}

func (s *Syntax) malformedError() {
	compileError("syntax-error: malformed %s", s.Bounder())
}

func (s *Syntax) assertListEqual(arguments Object, length int) {
	if !arguments.isList() || arguments.(*Pair).ListLength() != length {
		s.malformedError()
	}
}

func (s *Syntax) assertListMinimum(arguments Object, minimum int) {
	if !arguments.isList() || arguments.(*Pair).ListLength() < minimum {
		s.malformedError()
	}
}

func (s *Syntax) assertListRange(arguments Object, lengthRange []int) {
	if !arguments.isList() {
		s.malformedError()
	}

	for _, length := range lengthRange {
		if length == arguments.(*Pair).ListLength() {
			return
		}
	}
	s.malformedError()
}

func setSyntax(s *Syntax, arguments Object) Object {
	s.assertListEqual(arguments, 2)
	elements := arguments.(*Pair).Elements()

	variable := elements[0]
	if !variable.isVariable() {
		s.malformedError()
	}
	value := elements[1].Eval()
	s.Bounder().updateBinding(variable.(*Variable).identifier, value)
	return undef
}

func ifSyntax(s *Syntax, arguments Object) Object {
	s.assertListRange(arguments, []int{2, 3})
	elements := arguments.(*Pair).Elements()

	result := elements[0].Eval()
	if result.isBoolean() && !result.(*Boolean).value {
		if len(elements) == 3 {
			return elements[2].Eval()
		} else {
			return undef
		}
	} else {
		return elements[1].Eval()
	}
}

func andSyntax(s *Syntax, arguments Object) Object {
	s.assertListMinimum(arguments, 0)

	lastResult := Object(NewBoolean(true))
	for _, object := range arguments.(*Pair).Elements() {
		lastResult = object.Eval()
		if lastResult.isBoolean() && lastResult.(*Boolean).value == false {
			return NewBoolean(false)
		}
	}
	return lastResult
}

func orSyntax(s *Syntax, arguments Object) Object {
	s.assertListMinimum(arguments, 0)

	lastResult := Object(NewBoolean(false))
	for _, object := range arguments.(*Pair).Elements() {
		lastResult = object.Eval()
		if !lastResult.isBoolean() || lastResult.(*Boolean).value != false {
			return lastResult
		}
	}
	return lastResult
}

func beginSyntax(s *Syntax, arguments Object) Object {
	s.assertListMinimum(arguments, 0)

	lastResult := Object(undef)
	for _, object := range arguments.(*Pair).Elements() {
		lastResult = object.Eval()
	}
	return lastResult
}

func defineSyntax(s *Syntax, arguments Object) Object {
	s.assertListEqual(arguments, 2)
	elements := arguments.(*Pair).Elements()

	if !elements[0].isVariable() {
		runtimeError("Compile Error: syntax-error: (define)")
	}
	variable := elements[0].(*Variable)
	s.Bounder().bind(variable.identifier, elements[1].Eval())

	return NewSymbol(variable.identifier)
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
