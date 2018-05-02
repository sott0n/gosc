# go-scheme
[![CircleCI](https://circleci.com/gh/sott0n/go-scheme.svg?style=shield)](https://circleci.com/gh/sott0n/go-scheme)
[![](http://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

**go-scheme** is a toy scheme processing by Go-lang.

## Syntax and Function

| Type | To be done | Implemented |
|:-----|:-----|:-----------:|
| Number | number?, *, /, =, <, <=, >, >= | +, - |
| List | null?, pair?, list?, symbol?, car, cdr, cons, list, length, memq, last, append, set-car!, set-cdr! |  |
| Boolean | boolean?, not |  |
| String | string?, string-append, symbol->string, string->symbol, string->number, number->string |  |
| Procedure | procedure? |  |
| Comparison | eq?, neq?, equal? |  |
| Syntax | lambda, let, let*, letrec |  |
| Statement | if, cond, and, or, begin, do |  |
| Definition | set!, define, define-macro |  |
| Others | load |  |

## TODO

* Create AST method
* Add some tests
* Add [-*/] symbol
* etc ..