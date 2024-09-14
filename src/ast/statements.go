package ast

type BlockStmt struct {
	// { ... []Stmt }
	Body []Stmt
}

func (block BlockStmt) stmt() {
	// Implement statement evaluation for BlockStmt
}

type ExpressionStmt struct {
	Expression Expr
}

func (expr ExpressionStmt) stmt() {
	// Implement statement evaluation for ExpressiumStmt
}
