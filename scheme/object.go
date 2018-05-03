// Object and ObjectBase is an abstruct class for all scheme expression.
// A return value of a method which returns scheme object is Object.
// And ObjectBase has Object's implementation of String().

package scheme

// Object is an abstruct class for scheme object.
type Object interface {
	String() string
	IsNumber() bool
	IsPair() bool
	IsList() bool
	IsApplication() bool
}

// ObjectBase is an abstruct class for base scheme object.
type ObjectBase struct {
}

func (o *ObjectBase) String() string {
	return "This type's String() is not implemented yet."
}

// IsNumber is an interface function of number boolean.
func (o *ObjectBase) IsNumber() bool {
	return false
}

// IsPair is an interface function of pair boolean.
func (o *ObjectBase) IsPair() bool {
	return false
}

// IsList is an interface function of list boolean.
func (o *ObjectBase) IsList() bool {
	return false
}

// IsApplication is an interface function of application boolean.
func (o *ObjectBase) IsApplication() bool {
	return false
}
