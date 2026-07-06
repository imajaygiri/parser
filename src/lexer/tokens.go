package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	INDENTIFIER
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN
	ASSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS
	OR
	AND
	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	NULLISH_ASSIGNMENT
	// SLASH_EQUALS
	// STAR_EQUALS
	PLUS
	DASH
	SLASH
	STAR
	PERCENT
	// Reserved keyword
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	Kind  TokenKind
	Value string
}

func TokenKindString(kind TokenKind) string {
	kindsInfo := []string{
		"EOF",
		"NUMBER",
		"STRING",
		"INDENTIFIER",
		"OPEN_BRACKET",
		"CLOSE_BRACKET",
		"OPEN_CURLY",
		"CLOSE_CURLY",
		"OPEN_PAREN",
		"CLOSE_PAREN",
		"ASSIGNMENT",
		"EQUALS",
		"NOT",
		"NOT_EQUAL",
		"LESS",
		"LESS_EQUAL",
		"GREATER",
		"GREATER_EQUALS",
		"OR",
		"AND",
		"DOT",
		"DOT_DOT",
		"SEMI_COLON",
		"COLON",
		"QUESTION",
		"COMMA",
		"PLUS_PLUS",
		"MINUS_MINUS",
		"PLUS_EQUALS",
		"MINUS_EQUALS",
		"NULLISH_ASSIGNMENT",
		"PLUS",
		"DASH",
		"SLASH",
		"STAR",
		"PERCENT",
		"LET",
		"CONST",
		"CLASS",
		"NEW",
		"IMPORT",
		"FROM",
		"FN",
		"IF",
		"ELSE",
		"FOREACH",
		"WHILE",
		"FOR",
		"EXPORT",
		"TYPEOF",
		"IN",
	}
	return kindsInfo[kind]
}

func (kind TokenKind) ToString() string {
	return TokenKindString(kind)
}

func (token Token) Debug() {

	if token.isOneOfMany(NUMBER, INDENTIFIER, STRING) {
		fmt.Printf("%s %s", token.Kind.ToString(), token.Value)
	} else {
		fmt.Printf("%s", token.Kind.ToString())
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{Kind: kind, Value: value}
}

func (token Token) isOneOfMany(expectedToken ...TokenKind) bool {
	for _, expected := range expectedToken {
		if expected == token.Kind {
			return true
		}
	}
	return false
}
