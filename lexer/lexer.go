package lexer

import (
	"interpreter/token"
)

// readPosition is used as a peak ahead
// reads the char right after position to know
// what comes up next -> position corresponds
// to the character in ch byte
type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

// New creates a Lexer from
// an input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch

			l.readChar()

			tok.Type = token.NOT_EQ
			tok.Literal = string(l.ch) + string(ch)
		}
		tok = newToken(token.BANG, l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch

			l.readChar()

			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
		}
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// return is inserted here because inside of
		// readIdentifier, l.position and l.readPosition
		// are placed ahead of the last character on the identifier
		// so calling readChar again would place theses markers
		// at an incorrect position
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		}

		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// advances position and readPosition
// until it encounters a non letter char
// saves the original pos of the first letter char
// and crops the input so that the identifier is known
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	// EOF or next character
	switch {
	case l.readPosition >= len(l.input):
		l.ch = 0
	default:
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// does not advance the lexer's position markers
// it only peeks ahead and reports back the character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}