package diagnostics

import (
	"fmt"

	"github.com/kvexium/kvexc/src/diagnostics"
)

// FunctionRegistry speichert Funktionen anhand ihrer Namen und Argumenttypen
var FunctionRegistry = make(map[string]map[string]func(...interface{}))

// RegisterFunction registriert eine Funktion mit einem bestimmten Typ
func RegisterFunction(_type string, argType string, fn func(...interface{})) {
	if FunctionRegistry[_type] == nil {
		FunctionRegistry[_type] = make(map[string]func(...interface{}))
	}
	FunctionRegistry[_type][argType] = fn
}

// CallFunction ruft die richtige Funktion auf, basierend auf dem Funktionsnamen und dem Argumenttyp
func ThrowDiagnostic(_type string, argType string, args ...interface{}) error {
	if FunctionRegistry[_type] == nil {
		fmt.Printf("%s No matching diagnostics type '%s' found\n",
			diagnostics.FrontTextNamed(_type, string(diagnostics.Error), "DiagnosticsBag"), _type)
	}

	fn, ok := FunctionRegistry[_type][argType]
	if !ok {
		fmt.Printf("%s No matching argument type '%s' found\n",
			diagnostics.FrontTextNamed(_type, string(diagnostics.Error), "DiagnosticsBag"), argType)
	}

	fn(args...)
	return nil
}
