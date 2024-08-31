package parser

import (
	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

type bindingPower int

const (
	default_pb bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
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


func led(kind lexer.TokenKind, bp bindingPower, led_fn ledHandler) {
	bpLu[]
}