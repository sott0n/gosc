// renderIndentGuides

package scheme

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"text/scanner"
)

// Lexer is a struction for lexical analyzer.
type Lexer struct {
	scanner.Scanner
}

// EOF defined.
const (
	EOF = -(iota + 1)
	IdentifierToken
	IntToken
)

var identifierChars = "a-zA-Z?!*/<=>:$%^&_~"
var numberChars = "0-9+-."
var identifierExp = fmt.Sprintf("[%s][%s%s]*", identifierChars, identifierChars, numberChars)

// NewLexer is defining a new Lexer.
func NewLexer(source string) *Lexer {
	lexer := new(Lexer)
	lexer.Init(strings.NewReader(source))
	return lexer
}

// TokenType means Non-destructive scanner.Scan().
// This method returns next token type or unicode character.
func (l Lexer) TokenType() rune {
	token := l.PeekToken()
	if l.matchRegexp(token, "^[ ]*$") {
		return EOF
	} else if l.matchRegexp(token, fmt.Sprintf("^(%s|\\+|-)$", identifierExp)) {
		return IdentifierToken
	} else if l.matchRegexp(token, "^[0-9]+$") {
		return IntToken
	} else {
		runes := []rune(token)
		return runes[0]
	}
}

// PeekToken means Non-desructive Lexer.NextToken().
func (l Lexer) PeekToken() string {
	l.Scan()
	return l.TokenText()
}

// NextToken returns next token and moves current token reading
// position to next token position.
func (l *Lexer) NextToken() string {
	l.Scan()
	return l.TokenText()
}

// matchRegexp means matching token and string or symbol, integer..
func (l *Lexer) matchRegexp(matchString string, expression string) bool {
	re, err := regexp.Compile(expression)
	if err != nil {
		log.Fatal(err.Error())
	}
	return re.MatchString(matchString)
}
