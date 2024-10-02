// main.go
package main

import (
	"fmt"
	"os"

	"github.com/sanity-io/litter"

	"github.com/kvexium/kvexc/src/diagnostics"
	"github.com/kvexium/kvexc/src/lexer"
	"github.com/kvexium/kvexc/src/parser"
)

func main() {
	filePath := "./src/examples/"
	fileName := "03.kvex"

	currentDir, _ := os.Getwd()
	fmt.Printf("Current directory: %s\n", currentDir)

	diagbag := diagnostics.DiagnosticsBag

	// Versuche, die Datei zu lesen und gib einen Fehler aus, wenn dies nicht gelingt
	bytes, err := os.ReadFile(filePath + fileName)
	if err != nil {
		diagbag
	}

	// Gib den gelesenen Inhalt aus, um zu überprüfen, ob er korrekt eingelesen wurde
	fmt.Printf("File content:%s\n", string(bytes))

	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)

	litter.Dump(ast)
}
