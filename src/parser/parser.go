package parser

import (
	"fmt"

	"github.com/kvexium/kvexc/src/ast"
	"github.com/kvexium/kvexc/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
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

func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind
	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected '%s' but recieved '%s'", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}
	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
