package parser

import (
	"fmt"
	"strconv"

	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
	"github.com/imajaygiri/parser/src/utils"
)

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
