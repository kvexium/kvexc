package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func (lex *Lexer, regex *regexp.Regexp)

type regexPattern struct {
    regex *regexp.Regexp
	handler regexHandler
}

type Lexer struct {
	patterns []regexPattern
	Tokens []Token
    source string
    pos int
}

func (lex *Lexer) advanceN(n int) {
    lex.pos += n
}

func (lex *Lexer) push(token Token) {
    lex.Tokens = append(lex.Tokens, token)
}

func (lex *Lexer) at () byte {
    return lex.source[lex.pos]
}

func (lex *Lexer) remainder () string {
	return lex.source[lex.pos:]
}

func (lex *Lexer) at_eof () bool {
	return lex.pos >= len(lex.source)
}

func Tokenize (source string) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc == nil && loc[0] == 0 {
                matched = true
                pattern.handler(lex, pattern.regex)
            	break
			}
		}

		if !matched {
			// Error message for unrecognized token with line number and column number provided
			// ERROR [ Lexer ] -> Unrecognized token: <Token> at position <Line>:<Column>
			fmt.Printf("Error [ Lexer ] -> Unrecognized token: %s at position %d:%d\n", lex.remainder(), lex.pos, lex.pos - len(lex.remainder()) + 1)
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
        lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
    }
}

func createLexer(source string) *Lexer {
	return &Lexer{
		pos: 0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			// define all patterns
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\s+`), skipHandler},

			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
            {regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
            {regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
            {regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
            {regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
            {regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
            {regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
            {regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
            {regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
            {regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
            {regexp.MustCompile(`\.\.`), defaultHandler(DOUBLE_DOT, "..")},
            {regexp.MustCompile(`.`), defaultHandler(DOT, ".")},
            {regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
            {regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
            {regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
            {regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
            {regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
            {regexp.MustCompile(`\-\-`), defaultHandler(MINUS_MINUS, "--")},
            {regexp.MustCompile(`\-=`), defaultHandler(MINUS_EQUALS, "-=")},
            {regexp.MustCompile(`\-`), defaultHandler(MINUS, "-")},
			// {regex.MustCompile(`\*=`), defaultHandler(STAR_EQUALS, "*=")},
            {regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
            // {regex.MustCompile(`/=`), defaultHandler(SLASH_EQUALS, "/=")},
            {regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
            {regexp.MustCompile(`\%`), defaultHandler(MODULO, "%")},
		},
	}
}

func skipHandler(lex *Lexer, regex *regexp.Regexp) {
    match := regex.FindStringIndex(lex.remainder())
    lex.advanceN(match[1])
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}