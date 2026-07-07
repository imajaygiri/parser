package ast

// Example of block statemenet
// {... []Stmt}

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) Stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) Stmt() {
}
