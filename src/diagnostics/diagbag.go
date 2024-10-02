package diagnostics

import (
	"fmt"
	"strings"

	"github.com/fatih/color" // Verwende für farbige Ausgabe
)

type Type string

const (
	ErrorType   Type = "ERROR"
	WarningType Type = "WARNING"
	InfoType    Type = "INFO"
)

type Position struct {
	Line   int
	Column int
}

type Diagnostic struct {
	Type        Type
	File        string   // Dateiname
	Pos         Position // Position (Zeile, Spalte)
	Description string   // Beschreibung der Diagnose
	SourceLine  string   // Die Zeile des Quellcodes
}

type DiagnosticsBag struct {
	diagnostics []Diagnostic
}

// throwError - Fügt einen Error in die Diagnosen-Liste hinzu
func (bag *DiagnosticsBag) throwError(d Diagnostic) {
	d.Type = ErrorType
	bag.diagnostics = append(bag.diagnostics, d)
	bag.printDiagnostic(d)
}

// throwWarning - Fügt eine Warnung in die Diagnosen-Liste hinzu
func (bag *DiagnosticsBag) throwWarning(d Diagnostic) {
	d.Type = WarningType
	bag.diagnostics = append(bag.diagnostics, d)
	bag.printDiagnostic(d)
}

// throwInfo - Fügt eine Information in die Diagnosen-Liste hinzu
func (bag *DiagnosticsBag) throwInfo(d Diagnostic) {
	d.Type = InfoType
	bag.diagnostics = append(bag.diagnostics, d)
	bag.printDiagnostic(d)
}

// printDiagnostic - Gibt eine Diagnose im gewünschten Format aus
func (bag *DiagnosticsBag) printDiagnostic(d Diagnostic) {
	switch d.Type {
	case ErrorType:
		printError(d)
	case WarningType:
		printWarning(d)
	case InfoType:
		printInfo(d)
	}
}

// printError - Fehler mit Datei, Position und Caret anzeigen
func printError(d Diagnostic) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s: %s at %s:%d:%d\n", red("ERROR"), d.Description, d.File, d.Pos.Line, d.Pos.Column)
	fmt.Println(d.SourceLine)
	fmt.Println(generateCaret(d.Pos.Column))
}

// printWarning - Warnung mit Datei, Position und Caret anzeigen
func printWarning(d Diagnostic) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("%s: %s at %s:%d:%d\n", yellow("WARNING"), d.Description, d.File, d.Pos.Line, d.Pos.Column)
	fmt.Println(d.SourceLine)
	fmt.Println(generateCaret(d.Pos.Column))
}

// printInfo - Info mit Datei und Position anzeigen
func printInfo(d Diagnostic) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%s: %s at %s:%d:%d\n", blue("INFO"), d.Description, d.File, d.Pos.Line, d.Pos.Column)
	fmt.Println(d.SourceLine)
}

// Helper function to generate caret for error/warning
func generateCaret(column int) string {
	return fmt.Sprintf("%s^", strings.Repeat(" ", column-1))
}

// Beispiel zur Verwendung
func main() {
	bag := DiagnosticsBag{}

	// Beispiel-Fehler werfen
	bag.throwError(Diagnostic{
		File:        "main.kvx",
		Pos:         Position{Line: 10, Column: 5},
		Description: "Unexpected token '}'",
		SourceLine:  "let x = { 1, 2, };",
	})

	// Beispiel-Warnung werfen
	bag.throwWarning(Diagnostic{
		File:        "main.kvx",
		Pos:         Position{Line: 20, Column: 12},
		Description: "Unused variable 'y'",
		SourceLine:  "let y = 10;",
	})

	// Beispiel-Info werfen
	bag.throwInfo(Diagnostic{
		File:        "main.kvx",
		Pos:         Position{Line: 30, Column: 1},
		Description: "Compilation finished successfully",
		SourceLine:  "// end of file",
	})
}
