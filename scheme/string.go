// String is a type for scheme string object, which is
// expressed like "string".

package scheme

// String is a struction for scheme string object.
type String struct {
	ObjectBase
	text string
}

// NewString is a function for difinition a new String object.
func NewString(text string) *String {
	return &String{text: text}
}
