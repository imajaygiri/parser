package parser

import (
	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
)

// expr : 10 + 2 * 20
// stmt : let x = 10;
func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, ok := stmt_lu[p.currentTokenKind()]

	if ok {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}

}
