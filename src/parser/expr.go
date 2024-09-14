package parser

import (
	"fmt"
	"strconv"

	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	// First parse the NUD
	tokenkind := p.currentTokenKind()
	nudFn, exists := nudLu[tokenkind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for Token '%s'", lexer.TokenKindString(tokenkind)))
	}

	left := nudFn(p)

	for bpLu[p.currentTokenKind()] > bp {
		tokenKind := p.currentTokenKind()
		ledFn, exists := ledLu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED Handler expected Token '%s'", lexer.TokenKindString(tokenkind)))
		}

		left = ledFn(p, left, bp)
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUM:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{Value: number}
	case lexer.STR:
		return ast.StringExpr{Value: p.advance().Value}
	case lexer.IDENT:
		return ast.SymbolExpr{Value: p.advance().Value}
	default:
		panic(fmt.Sprintf("Cannot create Primary Expression from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operatorToken := p.advance()
	right := parseExpr(p, bp)

	return ast.BinaryExpr{
		Left:  left,
		Op:    operatorToken,
		Right: right,
	}
}
