package lexer

type TokenKind int

const {
	EOF TokenKind = iota
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

}

type Token struct {
	Kind TokenKind
	Value string
}