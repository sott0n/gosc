package scheme

// Definition is a type of define function.
type Definition struct {
	ObjectBase
	variable *Variable
	value    Object
}

// Eval is definition's eval function.
func (d *Definition) Eval() Object {
	d.bind(d.variable.identifier, d.value.Eval())
	return NewSymbol(d.variable.identifier)
}

func (d *Definition) String() string {
	return d.Eval().String()
}
