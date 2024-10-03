package diagnostics

import (
	"fmt"
)

// Colorcodes
type Colors string

const (
	Error   Colors = "#d77a61"
	Warning Colors = "#f2a65a"
	Info    Colors = "#52796f"
	Debug   Colors = "#d8b4a0"
)

// Beispiel für die Verwendung von RegisterFunction und CallFunction
func CreateAll() {
	FrontTextNamed := func(_type string, color string, name string) string {
		return fmt.Sprintf("[ %s :: %s ]", Colorize(color, []string{"bold"}, _type), name)
	}

	FrontText := func(_type string, color string) string {
		return fmt.Sprintf("[ %s ]", Colorize(color, []string{"bold"}, _type))
	}

	FrontTextPosNamed := func(_type string, color string, name string, file string, line int, column int) string {
		return fmt.Sprintf("[ %s :: %s | %s %d:%d ]", Colorize(color, []string{"bold"}, _type), name, file, line, column)
	}

	// DO NOT REMOVE !
	printDefaultArgsError := func(_type string, name string) {
		fmt.Println(FrontTextNamed("Error", string(Error), "DiagnosticsBag") + " Invalid number of arguments for [ " +
			_type + ", " + name + " ]")
	}

	// Funktion zur Fehlerausgabe definieren
	printArithmeticError := func(args ...interface{}) {
		_type := "Error"
		name := "Arithmetic"

		if len(args) != 6 {
			printDefaultArgsError(_type, name)
			return
		}
		file := args[0].(string)
		line := args[1].(int)
		column := args[2].(int)
		left := args[3].(string)
		right := args[4].(string)
		operation := args[5].(string)

		// 'type', Cannot <add> 'string'!
		fmt.Printf("%s '%s', Cannot %s '%s'\n",
			FrontTextPosNamed(_type, string(Error), name, file, line, column), left, operation, right)
	}

	printTokenError := func(args ...interface{}) {
		_type := "Error"
		name := "Token"

		if len(args) != 4 {
			printDefaultArgsError(_type, name)
			return
		}
		file := args[0].(string)
		line := args[1].(int)
		column := args[2].(int)
		expected := args[3].(string)

		fmt.Printf("%s Unexpected Token found! Expected Token '%s'\n",
			FrontTextPosNamed(_type, string(Error), name, file, line, column), expected)
	}

	printInputFileError := func(args ...interface{}) {
		_type := "Error"
		name := "InputFile"

		if len(args) != 2 {
			printDefaultArgsError(_type, name)
			return
		}
		file := args[0].(string)
		dir := args[1].(string)

		fmt.Printf("%s Cannot find file '%s' in directory '%s'\n",
			FrontTextNamed(_type, string(Error), name), file, dir)
	}

	printSourceCurrentDirDebug := func(args ...interface{}) {
		_type := "Debug"
		name := "SourceCurrentDirectory"

		if len(args) != 1 {
			printDefaultArgsError(_type, name)
			return
		}
		dir := args[0].(string)

		fmt.Printf("%s Current Directory: %s\n",
			FrontText(_type, string(Debug)), dir)
	}

	// Funktionen registrieren
	RegisterFunction("Error", "Arithmetic", printArithmeticError)
	RegisterFunction("Error", "Token", printTokenError)
	RegisterFunction("Error", "InputFile", printInputFileError)

	RegisterFunction("Debug", "SourceCurrentDirectory", printSourceCurrentDirDebug)
}
