// Application is a type to express scheme procedure application.
// Application has a procedure and its argument as list which consists.
// of Pair.

package scheme

import "log"

// Application is a struction for application.
type Application struct {
	ObjectBase
	procedureVariable Object
	arguments         Object // expect *Pair
	environment       *Environment
}

// Eval is eval function IF that returns applyProcedure.
func (a *Application) Eval() Object {
	return a.applyProcedure()
}

func (a *Application) String() string {
	result := a.Eval()
	if result == nil {
		log.Fatal("Procedure application returns nil.")
	}
	return result.String()
}

func (a *Application) applyProcedure() Object {
	if a.environment == nil {
		log.Fatal("Procedure does not have environment.")
	}
	return a.environment.invokeProcedure(a.procedureVariable, a.arguments)
}

// IsApplication is checking boolean wether is application or not.
func (a *Application) IsApplication() bool {
	return true
}
