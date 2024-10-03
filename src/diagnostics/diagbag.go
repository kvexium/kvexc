package diagnostics

import (
	"errors"
)

// FunctionRegistry speichert Funktionen anhand ihrer Namen und Argumenttypen
var FunctionRegistry = make(map[string]map[string]func(...interface{}))

// RegisterFunction registriert eine Funktion mit einem bestimmten Typ
func RegisterFunction(name string, argType string, fn func(...interface{})) {
	if FunctionRegistry[name] == nil {
		FunctionRegistry[name] = make(map[string]func(...interface{}))
	}
	FunctionRegistry[name][argType] = fn
}

// CallFunction ruft die richtige Funktion auf, basierend auf dem Funktionsnamen und dem Argumenttyp
func ThrowDiagnostic(name string, argType string, args ...interface{}) error {
	if FunctionRegistry[name] == nil {
		return errors.New("DiagnosticsBag: Diagnostic not found")
	}

	fn, ok := FunctionRegistry[name][argType]
	if !ok {
		return errors.New("DiagnosticsBag: No matching function found")
	}

	fn(args...)
	return nil
}
