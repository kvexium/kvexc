package lexer

import (
	"fmt"
	"regexp"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func Tokenize(source string) []Token {
	lex := createLexer(source)
	fmt.Printf("Source: %s\n", source) // Debug output

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			// Error message for unrecognized token with line number and column number provided
			// ERROR [ Lexer ] -> Unrecognized token: <Token> at position <Line>:<Column>
			fmt.Printf("ERROR [ Lexer ] -> Unrecognized token: %s at position %d:%d\n", lex.remainder(), lex.pos, lex.pos-len(lex.remainder())+1)
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

type regexHandler func(lex *lexer, regex *regexp.Regexp)

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		match := regex.FindString(lex.remainder())
		fmt.Printf("Matched: %s :: %s :: value: %s\n", match, TokenKindString(kind), value) // Debug output
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			// define all patterns
			// {regexp.MustCompile(`(\*(\*)?|&)?dec`), decHandler},
			{regexp.MustCompile(`(-)?[0-9]+(\.[0-9]+[fF])?`), numberHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`[a-z][a-zA-Z0-9]+`), symbolHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`\/\/.*`), commentHandler},
			{regexp.MustCompile(`\/\*(.|\s)*\*\/`), commentHandler},

			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGN, "=")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`\|`), defaultHandler(B_OR, "|")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`&`), defaultHandler(B_AND, "&")},
			{regexp.MustCompile(`\^`), defaultHandler(B_XOR, "^")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`::`), defaultHandler(IN, "::")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_ASSIGN, "+=")},
			{regexp.MustCompile(`\+\+`), defaultHandler(INCR, "++")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_ASSIGN, "--")},
			{regexp.MustCompile(`--`), defaultHandler(DECR, "--")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`~`), defaultHandler(MINUS, "~")},
			{regexp.MustCompile(`\*\*`), defaultHandler(EXPONENT, "**")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`\/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`%`), defaultHandler(MODULO, "%")},
			{regexp.MustCompile(`#`), defaultHandler(SW_VALUE, "#")},
			{regexp.MustCompile(`_`), defaultHandler(DEFAULT, "_")},
		},
	}
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	fmt.Printf("Integer Number Matched: %s\n", match) // Debug output

	isFloat := false

	if match[len(match)-1] == 'f' || match[len(match)-1] == 'F' {
		match = match[:len(match)-1]
		isFloat = true
	}

	// Erstelle ein Token mit der Ganzzahl
	lex.push(NewToken(NUM, match))

	// Bewege den Lexer vorwärts basierend auf der Länge des Matches
	if isFloat {
		lex.advanceN(len(match) + 1)
	} else {
		lex.advanceN(len(match))
	}
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]
	fmt.Printf("String Matched: %d\n", match) // Debug output
	lex.push(NewToken(STR, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

func commentHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	fmt.Printf("Symbol Matched: %s\n", match) // Debug output

	if kind, exists := reserved_keywords[match]; exists {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENT, match))
	}

	lex.advanceN(len(match))
}
