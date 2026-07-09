package parser

import (
	"fmt"
	"strconv"

	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
	"github.com/imajaygiri/parser/src/utils"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_handler, ok := nud_lu[tokenKind]
	if !ok {
		utils.Error(
			fmt.Sprintf("Couldn't find nud_handler for %s",
				tokenKind.ToString(),
			))
		panic(" -> Error finding nud_handler")
	}
	// pos of parser is updated by nud_handler
	left := nud_handler(p)

	for bp < bp_lu[p.currentTokenKind()] {
		tokenKind = p.currentTokenKind()
		led_handler, ok := led_lu[tokenKind]
		if !ok {
			utils.Error(
				fmt.Sprintf(
					"Couldn't find nud_handler for %s",
					tokenKind.ToString(),
				))
			panic(" -> Error finding led_handler")
		}
		left = led_handler(p, left, bp_lu[tokenKind])
	}
	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, err := strconv.ParseFloat(p.advance().Value, 64)

		if err != nil {
			panic(utils.Error("Error parsing number to float"))
		}
		return ast.NumberExpr{
			Value: number,
		}

	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}

	case lexer.INDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(utils.Error(fmt.Sprintf("can not create primary_expression from %s\n", p.currentTokenKind().ToString())))
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
