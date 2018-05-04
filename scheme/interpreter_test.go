package scheme

import (
	"testing"
)

type parserTest struct {
	source  string
	results []string
}

type evalErrorTest struct {
	source  string
	message string
}

var parserTests = []parserTest{
	makeTest("12", "12"),
	makeTest("()", "()"),
	makeTest("#f", "#f"),
	makeTest("#t", "#t"),
	makeTest("1234567890", "1234567890"),

	makeTest("\"\"", "\"\""),
	makeTest("\"hello\"", "\"hello\""),

	makeTest("(+)", "0"),
	makeTest("(- 1)", "1"),
	makeTest("(*)", "1"),
	makeTest("(/ 1)", "1"),

	makeTest("(+ 1 2)", "3"),
	makeTest("(+ 1 20 300 4000)", "4321"),
	makeTest("( + 1 2 3 )", "6"),
	makeTest("(+ 1 (+ 2 3) (+ 3 4))", "13"),
	makeTest("(- 3 (- 2 3) (+ 3 0))", "1"),
	makeTest("(* (* 3 3) 3)", "27"),
	makeTest("(/ 100 (/ 4 2))", "50"),
	makeTest("(+ (* 100 3) (/ (- 4 2) 2))", "301"),

	makeTest("(= 2 1)", "#f"),
	makeTest("(= (* 100 3) 300)", "#t"),

	makeTest("(not #f)", "#t"),
	makeTest("(not #t)", "#f"),
	makeTest("(not (number? ()))", "#t"),
	makeTest("(not 1)", "#f"),
	makeTest("(not ())", "#f"),

	makeTest("(car '(1))", "1"),
	makeTest("(cdr '(1))", "()"),
	makeTest("(car '(1 2))", "1"),
	makeTest("(cdr '(1 2))", "(2)"),

	makeTest("(list)", "()"),
	makeTest("(list 1 2 3)", "(1 2 3)"),
	makeTest("(cdr (list 1 2 3))", "(2 3)"),

	makeTest("(string-append)", "\"\""),
	makeTest("(string-append \"a\" \" \" \"b\")", "\"a b\""),

	makeTest("(string->symbol \"a\")", "a"),
	makeTest("(symbol->string 'a)", "\"a\""),

	makeTest("(number? 100)", "#t"),
	makeTest("(number? (+ 3 (* 2 8)))", "#t"),
	makeTest("(number? #t)", "#f"),
	makeTest("(number? ()", "#f"),

	makeTest("(procedure? 1)", "#f"),
	makeTest("(procedure? +)", "#t"),

	makeTest("(boolean? 1)", "#f"),
	makeTest("(boolean? ())", "#f"),
	makeTest("(boolean? #t)", "#t"),
	makeTest("(boolean? #f)", "#t"),
	makeTest("(boolean? (null? 1))", "#t"),

	makeTest("(pair? 1)", "#f"),
	makeTest("(pair? ())", "#f"),
	makeTest("(pair? '(1 2 3))", "#t"),

	makeTest("(list? 1)", "#f"),
	makeTest("(list? ())", "#t"),
	makeTest("(list? '(1 2 3))", "#t"),

	makeTest("(symbol? 1)", "#f"),
	makeTest("(symbol? 'hello)", "#t"),

	makeTest("(string? 1)", "#f"),
	makeTest("(string? \"\")", "#t"),
	makeTest("(string? \"hello\")", "#t"),

	makeTest("(null? 1)", "#f"),
	makeTest("(null? ()", "#t"),

	makeTest("(define x 1) x", "x", "1"),
	makeTest("(define x (+ 1 3)) x", "x", "4"),
	makeTest("(define x 1) (define y 2) (define z 3) (+ x (* y z))", "x", "y", "z", "7"),
	makeTest("(define x 1) (define x 2) x", "x", "x", "2"),

	makeTest("'12", "12"),
	makeTest("'hello", "hello"),
	makeTest("'#f", "#f"),
	makeTest("'#t", "#t"),
	makeTest("'(  1  2  3  )", "(1 2 3)"),
	makeTest("'( 1 ( 2 3 ) )", "(1 (2 3))"),

	makeTest("(quote 12)", "12"),
	makeTest("(quote hello)", "hello"),
	makeTest("(quote #f)", "#f"),
	makeTest("(quote #t)", "#t"),
	makeTest("(quote ( 1 (3) 4))", "(1 (3) 4)"),
}

var evalErrorTests = []evalErrorTest{
	{"(1)", "Invalid application."},
	{"hello", "Unbound variable: hello"},
	{"(quote)", "Compile Error: syntax-error: malformed quote."},
	{"(define)", "Compile Error: syntax-error: (define)"},

	{"(-)", "Compile Error: procedure requires at least 1 arguments."},
	{"(/)", "Compile Error: procedure requires at least 1 arguments."},
	{"(number?)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},

	{"(null?)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},
	{"(null? 1 2)", "Compile Error: Wrong number of arguments: number? requires 1, but got 2."},
	{"(not)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},

	{"(+ 1 #t)", "Compile Error: number required."},
	{"(- #t)", "Compile Error: number required."},
	{"(* ())", "Compile Error: number required."},
	{"(/ '(1 2 3))", "Compile Error: number required."},

	{"(string-append #f)", "Compile Error: string required."},
	{"(string-append 1)", "Compile Error: string required."},

	{"(string->symbol)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},
	{"(string->symbol 'hello)", "Compile Error: string required."},
	{"(symbol->string)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},
	{"(symbol->string \"\")", "Compile Error: symbol required."},

	{"(car ())", "Compile Error: pair required."},
	{"(cdr ())", "Compile Error: pair required."},
	{"(car)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},
	{"(cdr)", "Compile Error: Wrong number of arguments: number? requires 1, but got 0."},
}

func makeTest(source string, results ...string) parserTest {
	return parserTest{source: source, results: results}
}

func TestParser(t *testing.T) {
	for _, test := range parserTests {
		p := NewParser(test.source)
		p.Peek()
		for i := 0; i < len(test.results); i++ {
			result := test.results[i]
			parseObject := p.Parse()
			if parseObject == nil {
				t.Errorf("%s => <nil>; want %s", test.source, result)
				return
			}
			actual := parseObject.String()
			if actual != result {
				t.Errorf("%s => %s; want %s", test.source, actual, result)
			}
		}
	}
}

func TestEvalError(t *testing.T) {
	for _, test := range evalErrorTests {
		assertError(t, test.source, test.message)
	}
}

func assertError(t *testing.T, source string, message string) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("\"%s\" did not panic\n want: %s\n", source, message)
		} else if err != message {
			t.Errorf("\"%s\" paniced\nwith: %s\nwant: %s\n", source, err, message)
		}
	}()

	p := NewParser(source)
	p.Peek()
	p.Parse().Eval()
}
