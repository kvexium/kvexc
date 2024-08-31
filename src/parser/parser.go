package parser

import (
	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parseStmt(p))
	}

	return ast.BlockStmt{Body: Body}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().GetTokenKind()
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}
