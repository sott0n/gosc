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
	environment *Environment
	function    func(Object) Object
}

// NewProcedure is a function for definition a new procedure.
func NewProcedure(function func(Object) Object) *Procedure {
	return &Procedure{
		environment: nil,
		function:    function,
	}
}

// Eval is Procedure's eval IF.
func (p *Procedure) Eval() Object {
	return p
}

// Invoke is Procedure's function IF.
func (p *Procedure) Invoke(argument Object) Object {
	return p.function(argument)
}

// IsProcedure is boolean function IF.
func (p *Procedure) IsProcedure() bool {
	return true
}
