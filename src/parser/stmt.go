package parser

import (
	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

func parseStmt(p *parser) ast.Stmt {
	stmtFn, exists := stmtLu[p.currentTokenKind()]

	if exists {
		return stmtFn(p)
	}

	expression := parseExpr(p, defaultBp)
	p.expect(lexer.SEMICOLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}
