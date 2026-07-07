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

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"let":     LET,
	"const":   CONST,
	"class":   CLASS,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"fn":      FN,
	"if":      IF,
	"else":    ELSE,
	"foreach": FOREACH,
	"while":   WHILE,
	"for":     FOR,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"in":      IN,
}

func TokenKindString(kind TokenKind) string {
	kindsInfo := []string{
		"eof",
		"number",
		"string",
		"indentifier",
		"open_bracket",
		"close_bracket",
		"open_curly",
		"close_curly",
		"open_paren",
		"close_paren",
		"assignment",
		"equals",
		"not",
		"not_equal",
		"less",
		"less_equal",
		"greater",
		"greater_equals",
		"or",
		"and",
		"dot",
		"dot_dot",
		"semi_colon",
		"colon",
		"question",
		"comma",
		"plus_plus",
		"minus_minus",
		"plus_equals",
		"minus_equals",
		"nullish_assignment",
		"plus",
		"dash",
		"slash",
		"star",
		"percent",
		"let",
		"const",
		"class",
		"new",
		"import",
		"from",
		"fn",
		"if",
		"else",
		"foreach",
		"while",
		"for",
		"export",
		"typeof",
		"in",
	}
	return kindsInfo[kind]
}

func (kind TokenKind) ToString() string {
	return TokenKindString(kind)
}

func (token Token) Debug() {

	if token.isOneOfMany(NUMBER, INDENTIFIER, STRING) {
		fmt.Printf("%s (%s)", token.Kind.ToString(), token.Value)
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
