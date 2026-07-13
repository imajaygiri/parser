package parser

import (
	"fmt"
	"strconv"

	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/helpers"
	"github.com/imajaygiri/parser/src/lexer"
	"github.com/imajaygiri/parser/src/utils"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_handler, ok := nud_lu[tokenKind]
	if !ok {
		panic(fmt.Sprintf("No Nud handler for [%s]\n", tokenKind.ToString()))
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

// nud_handler
func parse_prefix_expr(p *parser) ast.Expr {
	// eg uninary(prefix_expr) -40;
	operatorToken := p.advance()
	rhs := parse_expr(p, default_bp)

	return ast.PrefixExpr{
		Operator:  operatorToken,
		RightExpr: rhs,
	}
}

// led_handler
func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, bp)
	return ast.AssignmentExpr{
		Operator: operatorToken,
		Value:    rhs,
		Assigne:  left, // to be assigned eg assigne = assigne + 1
	}
}

func parse_grouping_expr(p *parser) ast.Expr {
	p.advance() // grouping start (
	expr := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN) // exepected grouping end )
	return expr
}

func parse_struct_instantiation_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	structName := helpers.ExpectType[ast.SymbolExpr](left).Value
	properties := make(map[string]ast.Expr)
	p.expect(lexer.OPEN_CURLY)
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		varPropertyName := p.expect(lexer.INDENTIFIER).Value
		p.expect(lexer.COLON)
		expr := parse_expr(p, logical)

		_, exist := properties[varPropertyName]

		if exist {
			panic(fmt.Sprintf("Duplicate properties --> [%s]\n", varPropertyName))
		}
		properties[varPropertyName] = expr

		if p.currentTokenKind() != lexer.CLOSE_CURLY {
			p.expect(lexer.COMMA)
		}
	}
	p.expect(lexer.CLOSE_CURLY)

	return ast.StructInstantiation{
		StructName: structName,
		Properties: properties,
	}
}

func parse_array_instantiation_expr(p *parser) ast.Expr {
	//[]INDENTIFIER{1,2,3,4}
	var underlyingType ast.Type
	contents := make([]ast.Expr, 0)

	p.expect(lexer.OPEN_BRACKET)
	// array with size will be handled later [size]INDENTIFIER{}
	p.expect(lexer.CLOSE_BRACKET)
	underlyingType = parse_type(p, default_bp)

	p.expect(lexer.OPEN_CURLY)
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		contents = append(contents, parse_expr(p, logical))

		if p.currentTokenKind() != lexer.CLOSE_CURLY {
			p.expect(lexer.COMMA)
		}
	}
	p.expect(lexer.CLOSE_CURLY)

	return ast.ArrayInstantiationExpr{
		Underlying: underlyingType,
		Contents:   contents,
	}
}
