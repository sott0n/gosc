package scheme

import (
	"fmt"
	"regexp"
)

// Const is a scheme type.
type Const SchemeType

// NewConst is a new define const type.
func NewConst(expression string) *Const {
	return &Const{
		expression: expression,
	}
}

// String is a const of string function.
func (c *Const) String() string {
	return c.expression
}

// Eval is a const of eval function.
func (c *Const) Eval() (Type, error) {
	if matched, _ := regexp.MatchString("^[0-9]*$", c.expression); matched {
		return NewNumber(c.expression), nil
	} else if matched, _ := regexp.MatchString("^#(t|f)$", c.expression); matched {
		return NewBoolean(c.expression), nil
	} else {
		return nil, fmt.Errorf("Invalid or unexpected token: %s", c.expression)
	}
}
