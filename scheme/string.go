// String is a type for scheme string object, which is
// expressed like "string".

package scheme

import "fmt"

// String is a struction for scheme string object.
type String struct {
	ObjectBase
	text string
}

// NewString is a function for difinition a new String object.
func NewString(text string, options ...Object) *String {
	if len(options) > 0 {
		return &String{ObjectBase: ObjectBase{parent: options[0]}, text: text}
	}
	return &String{text: text}
}

// Eval is string's eval IF.
func (s *String) Eval() Object {
	return s
}

func (s *String) String() string {
	return fmt.Sprintf("\"%s\"", s.text)
}

func (s *String) isString() bool {
	return true
}
