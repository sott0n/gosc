// Symbol is a type to express scheme symbol object, which
// is expressed like 'symbol or (quote symbol).

package scheme

// Symbol is a struction for scheme symbol object.
type Symbol struct {
	ObjectBase
	identifier string
}

// NewSymbol is a function for difinition Symbol object.
func NewSymbol(identifier string) *Symbol {
	return &Symbol{identifier: identifier}
}

// Eval is symbol's eval IF.
func (s *Symbol) Eval() Object {
	return s
}

func (s *Symbol) String() string {
	return s.identifier
}

// IsSymbol returns boolean whethre symbol or not.
func (s *Symbol) IsSymbol() bool {
	return true
}
