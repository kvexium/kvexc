package diagnostics

import (
	"errors"
	"reflect"
)

// FunctionOverload definiert eine Struktur für eine überladene Funktion
type FunctionOverload struct {
	argType  string
	function reflect.Value
}

// Überladene Funktionen speichern
var OverloadedFunctions = make(map[string][]FunctionOverload)

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

	if len(overloads)%2 != 0 {
		fmt.Println("Invalid number of overload arguments.")
		return
	}

	for i := 0; i < len(overloads); i += 2 {
		argType, ok := overloads[i].(string)
		if !ok {
			fmt.Printf("Invalid argument type for overload %d\n", i)
			continue
		}

		function := reflect.ValueOf(overloads[i+1])
		if !function.IsValid() {
			fmt.Printf("Invalid function for overload %d\n", i)
			continue
		}

		functionOverloads = append(functionOverloads, FunctionOverload{
			argType:  argType,
			function: function,
		})
	}

	OverloadedFunctions[name] = functionOverloads
}


// addOverloadedFunction fügt eine neue überladene Funktion hinzu
func AddOverloadedFunction(name string, overloads ...interface{}) {
	for _, overload := range overloads {
		v := reflect.ValueOf(overload)
		OverloadedFunctions[name] = append(OverloadedFunctions[name], FunctionOverload{
			argType:  overload.(string), // Typ als String
			function: v,
		})
	}
}

// callOverloadedFunction ruft die richtige überladene Funktion auf
func CallOverloadedFunction(name string, argType string, args ...interface{}) error {
	overloads, exists := OverloadedFunctions[name]
	if !exists {
		return errors.New("Function not found")
	}

	for _, overload := range overloads {
		if argType != overload.argType {
			continue
		}

		overload.function.Call([]reflect.Value{reflect.ValueOf(args)})
		return nil
	}

	return errors.New("No matching function found")
}
