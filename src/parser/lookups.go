package parser

import (
	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
)

type binding_power int

const (
	default_bp     binding_power = iota // bp = 0
	comma                               // bp = 1
	assignment                          // bp = 2
	logical                             // bp = 3
	relational                          // bp = 4
	additive                            // bp = 5
	multiplicative                      // bp = 6
	unary                               // bp = 7
	call                                // bp = 8
	member                              // bp = 9
	primary                             // bp = 10
)

func (bp binding_power) Value() int {
	return int(bp)
}

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// Literals & symbols
	nud(lexer.NUMBER, parse_primary_expr)
	nud(lexer.STRING, parse_primary_expr)
	nud(lexer.INDENTIFIER, parse_primary_expr)
}
