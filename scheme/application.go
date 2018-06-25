// Application is a type to express scheme procedure application.
// Application has a procedure and its argument as list which consists.
// of Pair.

package scheme

// Application is a struction for application.
type Application struct {
	ObjectBase
	procedure Object
	arguments Object
}

// NewApplication is creating new object for application.
func NewApplication(parent Object) *Application {
	return &Application{
		ObjectBase: ObjectBase{parent: parent},
	}
}

// Eval is eval function IF that returns applyProcedure.
func (a *Application) Eval() Object {
	return a.applyProcedure()
}

func (a *Application) String() string {
	result := a.Eval()
	if result == nil {
		compileError("Procedure application returns nil")
	}
	return result.String()
}

func (a *Application) applyProcedure() Object {
	evaledObject := a.procedure.Eval()
	if !evaledObject.isProcedure() {
		runtimeError("invalid application")
	}
	procedure := evaledObject.(*Procedure)
	return procedure.Invoke(a.arguments)
}

func (a *Application) isApplication() bool {
	return true
}
