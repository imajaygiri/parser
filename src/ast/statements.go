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
	ExplicitType  Type   // future plan
	VariableName  string // let name = "ajay_giri";
	IsConstant    bool
	AssignedValue Expr
}

func (n VarDecStmt) stmt() {
}

type StructProperty struct {
	IsStatic bool // for later
	Type     Type
}

type StructMethod struct {
	IsStatic bool
	//Type FnType // for later
}

type StructDeclStmt struct {
	// Public     bool
	StructName string
	Properties map[string]StructProperty
	Methods     map[string]StructMethod
}

func (n StructDeclStmt) stmt() {
}
