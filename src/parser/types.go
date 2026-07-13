package parser

import (
	"fmt"
	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

func type_led(kind lexer.TokenKind, bp binding_power, type_led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = type_led_fn
}

func type_nud(kind lexer.TokenKind, type_nud_fn type_nud_handler) {
	type_nud_lu[kind] = type_nud_fn
}

func parse_type(p *parser, bp binding_power) ast.Type {
	tokenKind := p.currentTokenKind()
	type_nud_handler, ok := type_nud_lu[tokenKind]
	if !ok {
		panic(fmt.Sprintf("No type_nud_handler for [%s]\n", tokenKind.ToString()))
	}
	// pos of parser is updated by nud_handler
	left := type_nud_handler(p)

	for bp < type_bp_lu[p.currentTokenKind()] {
		tokenKind = p.currentTokenKind()
		type_led_handler, ok := type_led_lu[tokenKind]
		if !ok {

			panic(fmt.Sprintf(
				"Couldn't find type_led_handler for %s",
				tokenKind.ToString(),
			))
		}
		left = type_led_handler(p, left, type_bp_lu[tokenKind])
	}
	return left
}

func createTokenTypeLookup() {
	type_nud(lexer.INDENTIFIER, parse_symbol_type)
	type_nud(lexer.OPEN_BRACKET, parse_array_type)
}

func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.INDENTIFIER).Value,
	}
}

func parse_array_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CLOSE_BRACKET)
	underlyingType := parse_type(p, default_bp)
	return ast.ArrayType{
		Underlying: underlyingType,
	}
}
