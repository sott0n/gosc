package scheme

import (
	"testing"
)

func TestExpression(t *testing.T) {
	expectedResultByExpression := map[string]string{
		"1":            "1",
		"20":           "20",
		"300":          "300",
		"(+ 3 1)":      "4",
		"(number? 3)":  "#t",
		"(number? #t)": "#f",
		"(number? #f)": "#f",
	}

	for expression, expectResult := range expectedResultByExpression {
		actual := NewExpression(expression).String()

		if expectResult != actual {
			t.Errorf("Expected: %v, Got: %v", expectResult, actual)
			return
		}
	}

	invalidExpressions := []string{
		"(number? #t #t)",
	}

	for _, expression := range invalidExpressions {
		_, err := NewExpression(expression).Eval()
		if err == nil {
			t.Errorf("Expect %s to raise error, but got nil", expression)
		}
	}
}
