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

func (n ExpressionStmt) stmt() {
}

type VarDecStmt struct {
	VariableName  string // let name = "ajay_giri";
	IsConstant    bool
	AssignedValue Expr
	// ExplicitType Type // future plan
}

func (n VarDecStmt) stmt() {

}
