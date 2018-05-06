// Boolean is a type for scheme bool objects, such as #f and #t.

package scheme

// Boolean is a type object, true or false.
type Boolean struct {
	ObjectBase
	value bool
}

// NewBoolean is defining Boolean object.
func NewBoolean(value interface{}, options ...Object) (boolean *Boolean) {
	switch value.(type) {
	case bool:
		boolean = &Boolean{value: value.(bool)}
	case string:
		if value == "#t" {
			boolean = &Boolean{value: true}
		} else if value == "#f" {
			boolean = &Boolean{value: false}
		} else {
			compileError("Unexpected value for NewBoolean")
		}
	default:
		return nil
	}
	if len(options) > 0 {
		boolean.parent = options[0]
	}
	return
}

// Eval is boolean's eval IF.
func (b *Boolean) Eval() Object {
	return b
}

func (b *Boolean) String() string {
	if b.value == true {
		return "#t"
	}
	return "#f"
}

func (b *Boolean) isBoolean() bool {
	return true
}
