// Symbol is a type to express scheme symbol object, which
// is expressed like 'symbol or (quote symbol).

package scheme

var (
	symbols = make(map[string]*Symbol)
	undef   = Object(&Symbol{identifier: "#<undef>"})
)

// Symbol is a struction for scheme symbol object.
type Symbol struct {
	ObjectBase
	identifier string
}

// NewSymbol is a function for difinition Symbol object.
func NewSymbol(identifier string, options ...Object) *Symbol {
	if symbols[identifier] == nil {
		symbols[identifier] = &Symbol{ObjectBase: ObjectBase{parent: nil}, identifier: identifier}
	}
	return symbols[identifier]
}

// Eval is symbol's eval IF.
func (s *Symbol) Eval() Object {
	return s
}

func (s *Symbol) String() string {
	return s.identifier
}

func (s *Symbol) isSymbol() bool {
	return true
}
