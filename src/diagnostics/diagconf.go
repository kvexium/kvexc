package diagnostics

import (
	"fmt"
	"reflect"
)

// Funktion zur Fehlerausgabe definieren
func Create() {
	printArithmeticError := func(file string, line int, column int, token string, description string) {
		fmt.Printf("Arithmetic Error in %s at line %d, column %d: %s. Description: %s\n",
			file, line, column, token, description)
	}

	printTokenError := func(file string, line int, column int, token string) {
		fmt.Printf("Token Error in %s at line %d, column %d: %s\n",
			file, line, column, token)
	}

	// Überladene Funktionen registrieren
	CreateOverloadedFunction("Error",
		"Token", reflect.ValueOf(printTokenError),
		"Arithmetic", reflect.ValueOf(printArithmeticError),
	)

	// Füge eine neue überladene Funktion hinzu
	AddOverloadedFunction("Error",
		"Type", reflect.ValueOf(func(file string, line int, column int, token string, expected string) {
			fmt.Printf("Type Error in %s at line %d, column %d: %s. Expected: %s\n",
				file, line, column, token, expected)
		}),
		"Input", reflect.ValueOf(func(file string) {
			fmt.Printf("INput Error: Couldn't load file %s\n",
				file)
		}),
	)

	// Registriere einen benutzerdefinierten Fehler und rufe ihn auf
	/* myCustomError := func(message string) {
		fmt.Printf("Custom Error: %s\n", message)
	}

	createOverloadedFunction("CustomError", "Custom", reflect.ValueOf(myCustomError))
	callOverloadedFunction("CustomError", "Custom", "Something went wrong!") */
}
