package parser

import "github.com/kvexium/kvexc/src/lexer"

type parser struct {
	tokens []lexer.Token
	pos    int
}
