package ast

import "github.com/imajaygiri/parser/src/lexer"

// -------------------
// LITERRAL EXPRESSION
// -------------------
type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// -------------------
// COMPLEX EXPRESSION
// -------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}

type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (n PrefixExpr) expr() {}

// a = a + 5;
// a += 5;
// foo.bar += 10;
type AssignmentExpr struct {
	Assigne Expr
	Operator lexer.Token
	Value Expr
}

func (n AssignmentExpr) expr() {}
