package diagnostics

import "fmt"

// Beispiel für die Verwendung von RegisterFunction und CallFunction
func CreateAll() {
	// Funktion zur Fehlerausgabe definieren
	printArithmeticError := func(args ...interface{}) {
		if len(args) < 5 {
			fmt.Println("Invalid number of arguments")
			return
		}
		file := args[0].(string)
		line := args[1].(int)
		column := args[2].(int)
		token := args[3].(string)
		description := args[4].(string)

		fmt.Printf("Arithmetic Error in %s at line %d, column %d: %s. Description: %s\n",
			file, line, column, token, description)
	}

	printTokenError := func(args ...interface{}) {
		if len(args) < 4 {
			fmt.Println("Invalid number of arguments")
			return
		}
		file := args[0].(string)
		line := args[1].(int)
		column := args[2].(int)
		token := args[3].(string)

		fmt.Printf("Token Error in %s at line %d, column %d: %s\n",
			file, line, column, token)
	}

	printInputFileError := func(args ...interface{}) {
		if len(args) < 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		file := args[0].(string)

		fmt.Printf("Input File Error caused by: %s cannot be found\n",
			file)
	}

	// Funktionen registrieren
	RegisterFunction("Error", "Arithmetic", printArithmeticError)
	RegisterFunction("Error", "Token", printTokenError)
	RegisterFunction("Error", "InputFile", printInputFileError)
}
