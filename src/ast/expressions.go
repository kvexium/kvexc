package ast

import "github.com/kvexium/kvexc/src/lexer"

// -------------------
// LITERAL EXPRESSIONS
// -------------------

type NumberExpr struct {
	Value float64
}

func (num NumberExpr) expr() {
	// Implement expression evaluation for NumberExpr
}

type StringExpr struct {
	Value string
}

func (str StringExpr) expr() {
	// Implement expression evaluation for StringExpr
}

type SymbolExpr struct {
	Value string
}

func (sym SymbolExpr) expr() {
	// Implement expression evaluation for SymbolExpr
}

// -------------------
// COMPLEX EXPRESSIONS
// -------------------

type BinaryExpr struct {
	Left  Expr
	Op    lexer.Token
	Right Expr
}

func (bin BinaryExpr) expr() {
	// Implement expression evaluation for BinaryExpr
}
