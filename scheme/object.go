// Object and ObjectBase is an abstruct class for all scheme expression.
// A return value of a method which returns scheme object is Object.
// And ObjectBase has Object's implementation of String().

package scheme

// Object is an abstruct class for scheme object.
type Object interface {
	String() string
	IsProcedure() bool
}

// ObjectBase is an abstruct class for base scheme object.
type ObjectBase struct {
}

func (o *ObjectBase) String() string {
	return "This type's String() is not implemented yet."
}

// IsProcedure is boolean function whether it procedure or not.
func (o *ObjectBase) IsProcedure() bool {
	return false
}
