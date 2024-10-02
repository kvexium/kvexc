package parser

import (
	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

type bindingPower int

const (
	defaultBp bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	exponentiation
	unary
	call
	member
	primary
)

type stmtHandler func(p *parser) ast.Stmt
type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

type stmtLookup map[lexer.TokenKind]stmtHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

var bpLu = bpLookup{}
var nudLu = nudLookup{}
var ledLu = ledLookup{}
var stmtLu = stmtLookup{}

func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
	bpLu[kind] = bp
	ledLu[kind] = ledFn
}

func nud(kind lexer.TokenKind, nudFn nudHandler) {
	bpLu[kind] = primary
	nudLu[kind] = nudFn
}

func stmt(kind lexer.TokenKind, stmtFn stmtHandler) {
	bpLu[kind] = defaultBp
	stmtLu[kind] = stmtFn
}

func createTokenLookups() {
	// Logical
	led(lexer.AND, logical, parseBinaryExpr)
	led(lexer.OR, logical, parseBinaryExpr)

	// Relational
	led(lexer.EQUALS, relational, parseBinaryExpr)
	led(lexer.LESS, relational, parseBinaryExpr)
	led(lexer.LESS_EQUALS, relational, parseBinaryExpr)
	led(lexer.GREATER, relational, parseBinaryExpr)
	led(lexer.GREATER_EQUALS, relational, parseBinaryExpr)

	// Bitwise Relational
	led(lexer.B_AND, relational, parseBinaryExpr)
	led(lexer.B_XOR, relational, parseBinaryExpr)
	led(lexer.B_OR, relational, parseBinaryExpr)

	// TODO: IN Token

	// Additive & Multiplicative
	led(lexer.PLUS, additive, parseBinaryExpr)
	led(lexer.MINUS, additive, parseBinaryExpr)

	led(lexer.STAR, multiplicative, parseBinaryExpr)
	led(lexer.SLASH, multiplicative, parseBinaryExpr)
	led(lexer.MODULO, multiplicative, parseBinaryExpr)

	led(lexer.EXPONENT, exponentiation, parseBinaryExpr)

	/* // Unary
	nud(lexer.PLUS, unary, parseBinaryExpr)
	nud(lexer.MINUS, unary, parseBinaryExpr)
	nud(lexer.NOT, unary, parseBinaryExpr)
	nud(lexer.INCR, unary, parseBinaryExpr)
	nud(lexer.DECR, unary, parseBinaryExpr)
	nud(lexer.COMPL, unary parseBinaryExpr) */

	// Literals & Symbols
	nud(lexer.NUM, parsePrimaryExpr)
	nud(lexer.STR, parsePrimaryExpr)
	nud(lexer.IDENT, parsePrimaryExpr)
}
