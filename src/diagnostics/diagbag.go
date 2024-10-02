package diagnostics

import (
	"fmt"
	"reflect"
)

// FunctionOverload definiert eine Struktur für eine überladene Funktion
type FunctionOverload struct {
	argType  string
	function reflect.Value
}

// Überladene Funktionen speichern
var OverloadedFunctions = make(map[string][]FunctionOverload)

// CreateOverloadedFunction registriert eine überladene Funktion
func CreateOverloadedFunction(name string, overloads ...interface{}) {
	var functionOverloads []FunctionOverload

	for i := 0; i < len(overloads); i += 2 {
		argType := overloads[i].(string)            // Typ als String
		function := reflect.ValueOf(overloads[i+1]) // Funktion als Value
		functionOverloads = append(functionOverloads, FunctionOverload{
			argType:  argType,
			function: function,
		})
	}

	OverloadedFunctions[name] = functionOverloads
}

// CallOverloadedFunction ruft die richtige überladene Funktion auf
func CallOverloadedFunction(name string, argType string, args ...interface{}) {
	overloads, exists := OverloadedFunctions[name]
	if !exists {
		fmt.Println("Function not found")
		return
	}

	for _, overload := range overloads {
		if argType != overload.argType {
			continue
		}

		// Übergeben der Argumente in den richtigen Typen
		inArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			inArgs[i] = reflect.ValueOf(arg)
		}

		overload.function.Call(inArgs)
		return
	}

	fmt.Println("No matching function found")
}
