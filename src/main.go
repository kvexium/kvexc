package main

import (
	"fmt"
	"os"

	"github.com/kvexium/kvexc/src/diagnostics"
	"github.com/kvexium/kvexc/src/lexer"
	"github.com/kvexium/kvexc/src/parser"
	"github.com/kvexium/kvexc/src/registry"
	"github.com/sanity-io/litter"
)

func main() {
	registry.CreateAll()

	filePath := "./src/examples/"
	fileName := "02.kvex"

	currentDir, _ := os.Getwd()
	diagnostics.ThrowDiagnostic("Debug", "SourceCurrentDirectory", currentDir)

	// Versuche, die Datei zu lesen und gib einen Fehler aus, wenn dies nicht gelingt
	bytes, err := os.ReadFile(filePath + fileName)
	if err != nil {
		// Fehlerbehandlung: Aufruf der überladenen Funktion
		diagnostics.ThrowDiagnostic("Error", "InputFile", fileName, filePath)
		return // Beende die Funktion, um weitere Verarbeitung zu vermeiden
	}

	// Gib den gelesenen Inhalt aus, um zu überprüfen, ob er korrekt eingelesen wurde
	fmt.Printf("Dateiinhalt:\n%s\n", string(bytes))

	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)

	litter.Dump(ast)
}
