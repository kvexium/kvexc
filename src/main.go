// main.go
package main

import (
	"fmt"
	"os"

	"github.com/kvexium/kvexc/src/lexer"
)

func main() {
	filePath := "./src/examples/"
	fileName := "00.kvex"

	currentDir, _ := os.Getwd()
	fmt.Printf("Current directory: %s\n", currentDir)

	// Versuche, die Datei zu lesen und gib einen Fehler aus, wenn dies nicht gelingt
	bytes, err := os.ReadFile(filePath + fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	// Gib den gelesenen Inhalt aus, um zu überprüfen, ob er korrekt eingelesen wurde
	fmt.Printf("File content:%s\n", string(bytes))

	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
