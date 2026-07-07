package ast

// statement
type Stmt interface {
	stmt()
}

// expression: something which evaluates to somevalue
type Expr interface {
	expr()
}
