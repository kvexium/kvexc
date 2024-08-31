package ast

type BlockStmt struct {
	// { ... []Stmt }
	Body []Stmt
}

func (block BlockStmt) stmt() {
	// Implement statement evaluation for BlockStmt
}

type ExpressiumStmt struct {
	Expression Expr
}

func (expr ExpressiumStmt) stmt() {
	// Implement statement evaluation for ExpressiumStmt
}
