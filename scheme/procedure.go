// Procedure is a type for scheme procedure, which is expressed
// by lambda syntax form, like (lambda (x) x)
// when procedure has free variable, free variable must be binded when
// procedure is generated.
// So all Procedures have variable binding by Environment type (when there is
// no free variable, Procedure has Environmnet which is empty).

package scheme

// Procedure is a struction for scheme procedure.
type Procedure struct {
	ObjectBase
	function     func(Object) Object
	arguments    Object
	body         Object
	localBinding Binding
}

// NewProcedure is a function for definition a new procedure.
func NewProcedure(parent Object, arguments Object, body Object) *Procedure {
	// Create local binding for procedure.
	localBinding := parent.scopedBinding()
	procedure := &Procedure{
		ObjectBase:   ObjectBase{parent: nil},
		function:     nil,
		arguments:    arguments,
		body:         body,
		localBinding: localBinding,
	}

	procedure.function = func(givenArguments Object) Object {
		if !arguments.isList() || !givenArguments.isList() {
			runtimeError("Given non-list arguments")
		}

		// assert arguments count.
		expectedLength := arguments.(*Pair).ListLength()
		actualLength := givenArguments.(*Pair).ListLength()
		if expectedLength != actualLength {
			compileError("Wrong number of arguments: #f requires %d, but got %d",
				expectedLength, actualLength)
		}

		// binding arguments to local scope.
		parameters := arguments.(*Pair).Elements()
		objects := evaledObjects(givenArguments.(*Pair).Elements())
		for i, parameter := range parameters {
			if parameter.isVariable() {
				localBinding[parameter.(*Variable).identifier] = objects[i]
			}
		}

		// set procedure as a parent of body to eval elements in local binding
		body.setParent(procedure)

		// returns last eval result.
		var returnValue Object
		elements := body.(*Pair).Elements()
		for _, element := range elements {
			returnValue = element.Eval()
		}
		return returnValue
	}
	return procedure
}

func (p *Procedure) String() string {
	return "#<closure #f>"
}

// Eval is Procedure's eval IF.
func (p *Procedure) Eval() Object {
	return p
}

// Invoke is Procedure's function IF.
func (p *Procedure) Invoke(argument Object) Object {
	return p.function(argument)
}

func (p *Procedure) isProcedure() bool {
	return true
}

func (p *Procedure) binding() Binding {
	return p.localBinding
}
