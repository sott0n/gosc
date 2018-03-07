package scheme

import (
	"testing"
)

func TestExpression(t *testing.T) {
	expectedResultByExpression := map[string]string{
		"1":       "1",
		"20":      "20",
		"300":     "300",
		"(+ 3 1)": "4",
	}

	for expression, expectResult := range expectedResultByExpression {
		actual := NewExpression(expression).String()

		if expectResult != actual {
			t.Errorf("Expected: %v, Got: %v", expectResult, actual)
			return
		}
	}
}
