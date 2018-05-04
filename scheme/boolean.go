// Boolean is a type for scheme bool objects, such as #f and #t.

package scheme

// Boolean is a type object, true or false.
type Boolean struct {
	ObjectBase
	value bool
}

// NewBoolean is defining Boolean object.
func NewBoolean(value interface{}) *Boolean {
	switch value.(type) {
	case bool:
		return &Boolean{value: value.(bool)}
	case string:
		if value == "#t" {
			return &Boolean{value: true}
		} else if value == "#f" {
			return &Boolean{value: false}
		} else {
			panic("Unexpected value for NewBoolean.")
		}
	}
	return nil
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

// IsBoolean is boolean IF.
func (b *Boolean) IsBoolean() bool {
	return true
}
