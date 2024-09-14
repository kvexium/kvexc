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

	NUM
	IDENT

	OPEN_BRACKET  // [
	CLOSE_BRACKET // ]
	OPEN_CURLY    // {
	CLOSE_CURLY   // }
	OPEN_PAREN    // (
	CLOSE_PAREN   // )

	ASSIGN       // =
	NOT          // !
	PLUS         // +
	PLUS_ASSIGN  // +=
	INCR         // ++
	MINUS        // -
	MINUS_ASSIGN // -=
	DECR         // --

	STAR     // *
	EXPONENT // **
	SLASH    // /
	MODULO   // %

	EQUALS         // ==
	NOT_EQUALS     //!=
	LESS           // <
	LESS_EQUALS    // <=
	GREATER        // >
	GREATER_EQUALS // >=
	B_AND          // &
	AND            // &&
	B_XOR          // ^
	B_OR           // |
	OR             // ||

	DOT       // .
	SEMICOLON // ;
	COLON     // :
	IN        // ::
	COMMA     // ,
	QUESTION  // ?

	APOSTROPHE // '
	QUOTE      // "

	// Reserved Keywords
	CLASS // class ...
	NEW
	USE // uses ...

	IF       // if (...) {...}
	WHILE    // while (...) {...}
	FOR      // for (...) {...}
	SWITCH   // switch (...) {...}
	DEFAULT  // _
	SW_VALUE // #

	// EXPORT
	// TYPEOF

	DEC // dec

	I8
	I16
	I32
	I64

	U8
	U16
	U32
	U64

	F32
	F64
	F80

	C64
	C128

	STR
	CHAR

	BOOL
)

var reserved_keywords map[string]TokenKind = map[string]TokenKind{
	"null":  NULL,
	"true":  TRUE,
	"false": FALSE,

	"new":   NEW,
	"uses":  USE,
	"class": CLASS,

	"if":     IF,
	"while":  WHILE,
	"for":    FOR,
	"switch": SWITCH,

	// "export":  EXPORT,
	// "typeof":  TYPEOF,

	"dec": DEC,

	"i8":  I8,
	"i16": I16,
	"i32": I32,
	"i64": I64,

	"u8":  U8,
	"u16": U16,
	"u32": U32,
	"u64": U64,

	"f32": F32,
	"f64": F64,
	"f80": F80,

	"c64":  C64,
	"c128": C128,

	"str":  STR,
	"char": CHAR,

	"bool": BOOL,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if token.Kind == expected {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.Kind == IDENT || token.Kind == NUM || token.Kind == STR {
		fmt.Printf("%s [ %s ]\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s: '%s'\n", TokenKindString(token.Kind), token.Value)
	}
}

// gettokenKind methode as TokenKind
func (token Token) GetTokenKind() TokenKind {
	return token.Kind
}

func NewToken(Kind TokenKind, value string) Token {
	return Token{
		Kind, value,
	}
}

// Map to store the string representations of TokenKind constants.
var KindToStringMap = map[TokenKind]string{
	EOF: "EOF",

	NULL: "NULL",

	OPEN_BRACKET:  "OPEN_BRACKET",
	CLOSE_BRACKET: "CLOSE_BRACKET",
	OPEN_CURLY:    "OPEN_CURLY",
	CLOSE_CURLY:   "CLOSE_CURLY",
	OPEN_PAREN:    "OPEN_PAREN",
	CLOSE_PAREN:   "CLOSE_PAREN",

	ASSIGN:         "ASSIGN",
	NOT:            "NOT",
	LESS:           "LESS",
	LESS_EQUALS:    "LESS_EQUAL",
	GREATER:        "GREATER",
	GREATER_EQUALS: "GREATER_EQUAL",
	B_AND:          "B_AND",
	AND:            "AND",
	B_OR:           "B_OR",
	OR:             "OR",

	DOT:       "DOT",
	SEMICOLON: "SEMICOLON",
	COLON:     "COLON",
	IN:        "IN",
	QUESTION:  "QUESTION",
	COMMA:     "COMMA",

	APOSTROPHE: "APOSTROPHE",
	QUOTE:      "QUOTE",

	EQUALS:       "EQUALS",
	PLUS:         "PLUS",
	PLUS_ASSIGN:  "PLUS_ASSIGN",
	INCR:         "INCR",
	MINUS:        "MINUS",
	MINUS_ASSIGN: "MINUS_ASSIGN",
	DECR:         "DECR",
	STAR:         "STAR",
	EXPONENT:     "EXPONENT",
	SLASH:        "SLASH",
	MODULO:       "MODULO",

	TRUE:  "TRUE",
	FALSE: "FALSE",

	CLASS: "CLASS",
	NEW:   "NEW",
	USE:   "USE",

	IF:    "IF",
	WHILE: "WHILE",
	FOR:   "FOR",

	// EXPORT:         "EXPORT",
	// TYPEOF:         "TYPEOF",

	DEC: "DEC",

	I8:  "I8",
	I16: "I16",
	I32: "I32",
	I64: "I64",

	U8:  "U8",
	U16: "U16",
	U32: "U32",
	U64: "U64",

	F32: "F32",
	F64: "F64",
	F80: "F80",

	C64:  "C64",
	C128: "C128",

	BOOL: "BOOL",
	NUM:  "NUM",

	CHAR:  "CHAR",
	STR:   "STR",
	IDENT: "IDENT",
}

// TokenKindString returns the string representation of a TokenKind.
func TokenKindString(Kind TokenKind) string {
	if str, ok := KindToStringMap[Kind]; ok {
		return str
	}
	return "UNKNOWN"
}
