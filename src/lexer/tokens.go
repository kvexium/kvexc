package lexer

import (
	"fmt"
)

type TokenKind int

const (
	EOF TokenKind = iota

	NULL
	TRUE
	FALSE

	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT // =
	EQUALS     // ==
	NOT        // !
	NOT_EQUALS // !=

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS
	AND
	OR

	DOT
	DOUBLE_DOT
	SEMICOLON
	COLON
	QUESTION
	COMMA

	PLUS
	PLUS_PLUS
	PLUS_EQUALS
	MINUS
	MINUS_MINUS
	MINUS_EQUALS

	STAR
	// STAR_EQUALS
	SLASH
	// SLASH_EQUALS
	MODULO

	// Reserved Keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FUNC
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	kind TokenKind
	value string
}

func (token Token) isOneOfMany (expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
        if token.kind == expected {
            return true
        }
    }
    return false
}

func (token Token) Debug () {
	if token.kind == IDENTIFIER || token.kind == NUMBER || token.kind == STRING {
		fmt.Printf("%s (%s)\n", TokenKindString(token.kind), token.value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}

// Map to store the string representations of TokenKind constants.
var kindToStringMap = map[TokenKind]string{
	EOF:            "EOF",
	NULL:           "NULL",
	NUMBER:         "NUMBER",
	STRING:         "STRING",
	IDENTIFIER:     "IDENTIFIER",
	OPEN_BRACKET:   "OPEN_BRACKET",
	CLOSE_BRACKET:  "CLOSE_BRACKET",
	OPEN_CURLY:     "OPEN_CURLY",
	CLOSE_CURLY:    "CLOSE_CURLY",
	OPEN_PAREN:     "OPEN_PAREN",
	CLOSE_PAREN:    "CLOSE_PAREN",
	ASSIGNMENT:     "ASSIGNMENT",
	EQUALS:         "EQUALS",
	NOT:            "NOT",
	NOT_EQUALS:     "NOT_EQUALS",
	LESS:           "LESS",
	LESS_EQUALS:    "LESS_EQUALS",
	GREATER:        "GREATER",
	GREATER_EQUALS: "GREATER_EQUALS",
	AND:            "AND",
    OR:             "OR",
	DOT:            "DOT",
	DOUBLE_DOT:     "DOUBLE_DOT",
	SEMICOLON:      "SEMICOLON",
	COLON:          "COLON",
	QUESTION:       "QUESTION",
	COMMA:          "COMMA",
	PLUS:           "PLUS",
	PLUS_PLUS:      "PLUS_PLUS",
	PLUS_EQUALS:    "PLUS_EQUALS",
	MINUS:          "MINUS",
	MINUS_MINUS:    "MINUS_MINUS",
	MINUS_EQUALS:   "MINUS_EQUALS",
	STAR:           "STAR",
	SLASH:          "SLASH",
	MODULO:         "MODULO",
	NEW:            "NEW",
	CONST:          "CONST",
	CLASS:          "CLASS",
	IMPORT:         "IMPORT",
	FROM:           "FROM",
	FUNC:           "FUNC",
	IF:             "IF",
	ELSE:           "ELSE",
	FOREACH:        "FOREACH",
	WHILE:          "WHILE",
	FOR:            "FOR",
	EXPORT:         "EXPORT",
	TYPEOF:         "TYPEOF",
	IN:             "IN",
}

// TokenKindString returns the string representation of a TokenKind.
func TokenKindString(kind TokenKind) string {
	if str, ok := kindToStringMap[kind]; ok {
		return str
	}
	return "UNKNOWN"
}