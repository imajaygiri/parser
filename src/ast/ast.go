package ast

// statement
type Stmt interface {
	stmt()
}

// expression: something which evaluates to somevalue
type Expr interface {
	expr()
}

type Type interface {
	_type()
}

type StructInstantiation struct {
	StructName string
	Properties map[string]Expr
}

func (n StructInstantiation) expr() {}

type ArrayInstantiationExpr struct {
	Underlying Type
	Contents   []Expr
}

func (n ArrayInstantiationExpr) expr(){}
