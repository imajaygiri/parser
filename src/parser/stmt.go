package parser

import (
	"fmt"

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

func parse_var_dec_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	var assignedValue ast.Expr
	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(
		p.currentTokenKind(),
		"Expected variableName[indentifier]\n",
	).Value

	// check if colon is present for explicitType
	if p.currentTokenKind() == lexer.COLON {
		// consume colon
		p.advance()
		explicitType = parse_type(p, default_bp)
	}

	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Missing either right-hand-side in var declaration or type anotation.")
	}

	if isConstant && assignedValue == nil {
		panic("Cannot declare const without value.")
	}

	p.expect(lexer.SEMI_COLON)

	return ast.VarDecStmt{
		ExplicitType:  explicitType,
		VariableName:  varName,
		IsConstant:    isConstant,
		AssignedValue: assignedValue,
	}
}

func parse_struct_decl_stmt(p *parser) ast.Stmt {
	p.expect(lexer.STRUCT)
	properties := make(map[string]ast.StructProperty)
	methods := make(map[string]ast.StructMethod)
	structName := p.expect(lexer.INDENTIFIER).Value

	p.expect(lexer.OPEN_CURLY)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		var isStatic bool

		if p.currentTokenKind() == lexer.STATIC {
			isStatic = true
			p.expect(lexer.STATIC)
		}

		if p.currentTokenKind() == lexer.INDENTIFIER {
			propertyName := p.expect(lexer.INDENTIFIER).Value
			p.expectError(
				lexer.COLON,
				fmt.Sprintf("Expected --> [SEMI_COLON] found --> [%v]",
					p.currentTokenKind().ToString()),
			)
			strucType := parse_type(p, default_bp)
			p.expect(lexer.SEMI_COLON)

			_, exist := properties[propertyName]
			if exist {
				panic(fmt.Sprintf(
					"Duplicate properties --> [%v] found inside struct.",
					propertyName,
				))
			}

			properties[propertyName] = ast.StructProperty{
				IsStatic: isStatic,
				Type:     strucType,
			}

			continue
		}

		// methods
		// if p.currentTokenKind() == lexer.FN {
		// 	p.expect(lexer.FN) // eat fn token
		// 	// call method handler
		// }
		panic("Methods is not handled currently inside struct_vec_decl\n")
	}

	p.expect(lexer.CLOSE_CURLY)
	return ast.StructDeclStmt{
		StructName: structName,
		Properties: properties,
		Methods:    methods,
	}

}
