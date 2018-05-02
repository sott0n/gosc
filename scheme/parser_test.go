package scheme

import (
	"testing"
)

type parserTest struct {
	source string
	result string
}

var parserTests = []parserTest{
	{"()", "()"},
	{"#f", "#f"},
	{"#t", "#t"},
	{"1234567890", "1234567890"},
	{"(+)", "0"},
	{"(- 1)", "1"},
	{"(*)", "1"},
	{"(/ 1)", "1"},
	{"(+ 1 2)", "3"},
	{"(+ 1 20 300 4000)", "4321"},
	{"( + 1 2 3 )", "6"},
	{"(+ 1 (+ 2 3) (+ 3 4))", "13"},
	{"(- 3 (- 2 3) (+ 3 0))", "1"},
	{"(* (* 3 3) 3)", "27"},
	{"(/ 100 (/ 4 2))", "50"},
	{"(+ (* 100 3) (/ (- 4 2) 2))", "301"},
	{"(number? 100)", "#t"},
	{"(number? (+ 3 (* 2 8)))", "#t"},
	{"(number? #t)", "#f"},
	{"(number? ()", "#f"},
}

func TestParser(t *testing.T) {
	for _, test := range parserTests {
		p := NewParser(test.source)
		p.Peek()
		parseObject := p.Parse()
		if parseObject == nil {
			t.Errorf("%s => <nil>; want %s", test.source, test.result)
			return
		}
		actual := parseObject.String()
		if actual != test.result {
			t.Errorf("%s => %s; want %s", test.source, actual, test.result)
		}
	}
}
