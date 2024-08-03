package main

import (
    "os"

    "github.com/tlaceby/parser-series/src/lexer"
)

func main () {
	bytes, _ := os.ReadFile("./examples/00.kvxm")
	// source := string(bytes)

	tokens := lexer.Tokenize(string(bytes))
	// fmt.Printf("Code: {%s}\n", source)

	for _, token := range tokens {
        token.Debug()
    }
}