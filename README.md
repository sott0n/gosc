# gosc
[![CircleCI](https://circleci.com/gh/sott0n/gosc.svg?style=shield)](https://circleci.com/gh/sott0n/gosc)
[![](http://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

**gosc** is a toy scheme processing by Golang.

## How to install

You only type `go get` below:  

```bash
$ go get github.com/sott0n/gosc
```

## How to use

#### REPL mode : invoke interactive shell

```bash
$ gosc
```

#### Execute scheme file

```bash
$ gosc -f [filename].scm
```

#### One liner

```bash
$ gosc -e "(+ 1 2)"
```

#### Dump AST of input source code

```bash
$ gosc -a
```

#### Show Help

```bash
$ gosc -h
```

## Syntax and Function

| Type | Support | Status |
|:-|:-|:-:|
| Number | +, -, *, /, =, <, <=, >, >= | ○ |
| List | car, cdr, cons, list, length, memq, last, append, set-car!, set-cdr! | △ |
| Boolean | not, #f, #t | ○ |
| String | string-append, symbol->string, string->symbol, string->number, number->string | ○ |
| Type | number?, null?, pair?, list?, symbol?, procedure?, boolean?, string? | ○ |
| Comparison | eq?, neq?, equal? | ○ |
| Syntax | lambda, let, let*, letrec | △ |
| Statement | if, cond, and, or, begin, do | × |
| Definition | set!, define, define-macro | △ |
| Others | load | × |

## TODO

* Create AST method
* Add some tests
* Add [-*/] symbol
* etc ..