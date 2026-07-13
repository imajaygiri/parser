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
	// bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// assignment
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)

	//TODO:  *= , /= %=

	// logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	led(lexer.DOT_DOT, logical, parse_binary_expr) // 10..math.random()
	// relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)
	//additive and multiplicative
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)
	// Literals & symbols
	nud(lexer.NUMBER, parse_primary_expr)
	nud(lexer.STRING, parse_primary_expr)
	nud(lexer.INDENTIFIER, parse_primary_expr)
	nud(lexer.DASH, parse_prefix_expr)
	nud(lexer.OPEN_PAREN, parse_grouping_expr)

	//call/memeber/arrayInstantition expxr
	led(lexer.OPEN_CURLY, call, parse_struct_instantiation_expr)
	nud(lexer.OPEN_BRACKET, parse_array_instantiation_expr)

	//statements
	stmt(lexer.CONST, parse_var_dec_stmt)
	stmt(lexer.LET, parse_var_dec_stmt)
	stmt(lexer.STRUCT, parse_struct_decl_stmt)
}
