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
	BooleanToken
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
	} else if l.matchRegexp(token, "^#(f|t)$") {
		return BooleanToken
	} else {
		runes := []rune(token)
		return runes[0]
	}
}

// PeekToken means Non-desructive Lexer.NextToken().
func (l Lexer) PeekToken() string {
	return l.NextToken()
}

// NextToken returns next token and moves current token reading
// position to next token position.
func (l *Lexer) NextToken() string {
	return l.nextToken()
}

// IndentLevel return position of indent from symbol ( and ).
func (l *Lexer) IndentLevel() int {
	tokens := l.AllTokens()
	openCount, closedCount := 0, 0
	for _, token := range tokens {
		if token == "(" {
			openCount++
		} else if token == ")" {
			closedCount++
		}
	}
	return openCount - closedCount
}

// AllTokens returns token's list.
func (l *Lexer) AllTokens() []string {
	tokens := []string{}
	for {
		token := l.NextToken()
		if token == "" {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func (l *Lexer) nextToken() string {
	l.Scan()
	if l.TokenText() == "#" {
		// text/scanner scans '#t' as '#' and 't'.
		l.Scan()
		switch l.TokenText() {
		case "t", "f":
			return fmt.Sprintf("#%s", l.TokenText())
		default:
			log.Fatal("Tokens which start from '#' are not implemented except #f, #t.")
		}
	} else if l.matchRegexp(l.TokenText(), fmt.Sprintf("^%s$", identifierExp)) {
		// text/scanner scans some signs as splitted token from alphabet token.
		text := l.TokenText()
		for l.isIdentifierChar(l.Peek()) {
			l.Scan()
			text = fmt.Sprintf("%s%s", text, l.TokenText())
		}
		return text
	}
	return l.TokenText()
}

func (l *Lexer) isIdentifierChar(char rune) bool {
	charString := fmt.Sprintf("%c", char)
	return l.matchRegexp(charString, fmt.Sprintf("^[%s%s]$", identifierChars, numberChars))
}

// matchRegexp means matching token and string or symbol, integer..
func (l *Lexer) matchRegexp(matchString string, expression string) bool {
	re, err := regexp.Compile(expression)
	if err != nil {
		log.Fatal(err.Error())
	}
	return re.MatchString(matchString)
}
