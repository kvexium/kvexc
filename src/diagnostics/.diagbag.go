// package diagnostics

import (
	"fmt"
	"path/filepath"
	"strings"
)

type DiagnosticType string

const (
	ErrorType   DiagnosticType = "ERROR"
	WarningType DiagnosticType = "WARNING"
	InfoType    DiagnosticType = "INFO"
)

type MessageType int

const (
	MainMessage MessageType = iota
	Tip
	Note
)

type Message struct {
	Type    MessageType
	Content string
}

type Diagnostic struct {
	Type        DiagnosticType
	FilePath    string
	Messages    []Message
	LineNumber  int
	LineOfCode  string
	StartColumn int
	EndColumn   int
}

type DiagnosticsBag struct {
	diagnostics []Diagnostic
}

func NewDiagnosticsBag() *DiagnosticsBag {
	return &DiagnosticsBag{
		diagnostics: []Diagnostic{},
	}
}

func (db *DiagnosticsBag) AddDiagnostic(diagType DiagnosticType, filePath string, messages []Message, lineNumber int, lineOfCode string, startColumn, endColumn int) {
	diagnostic := Diagnostic{
		Type:        diagType,
		FilePath:    filePath,
		Messages:    messages,
		LineNumber:  lineNumber,
		LineOfCode:  lineOfCode,
		StartColumn: startColumn,
		EndColumn:   endColumn,
	}
	db.diagnostics = append(db.diagnostics, diagnostic)
}

func (db *DiagnosticsBag) HasErrors() bool {
	for _, d := range db.diagnostics {
		if d.Type == ErrorType {
			return true
		}
	}
	return false
}

func colorize(text string, color string) string {
	colorCodes := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"reset":  "\033[0m",
	}
	return colorCodes[color] + text + colorCodes["reset"]
}

func (db *DiagnosticsBag) String() string {
	var builder strings.Builder

	for i, d := range db.diagnostics {
		relPath, err := filepath.Rel(".", d.FilePath)
		if err != nil {
			relPath = d.FilePath
		}

		typeColor := "red"
		if d.Type == WarningType {
			typeColor = "yellow"
		} else if d.Type == InfoType {
			typeColor = "blue"
		}

		fmt.Fprintf(&builder, "/ -- [ %s ] :: %s\n", colorize(string(d.Type), typeColor), relPath)
		fmt.Fprintf(&builder, "|\n")

		maxLineNumWidth := len(fmt.Sprintf("%d", d.LineNumber))

		for _, msg := range d.Messages {
			var prefix string
			switch msg.Type {
			case MainMessage:
				prefix = "/ --"
			case Tip:
				prefix = "+ --"
			case Note:
				prefix = "= --"
			}
			fmt.Fprintf(&builder, "|%s %s\n", strings.Repeat(" ", maxLineNumWidth+1), colorize(prefix, "green"))
			fmt.Fprintf(&builder, "|%s %s\n", strings.Repeat(" ", maxLineNumWidth+1), msg.Content)
		}

		fmt.Fprintf(&builder, "| %*d | %s\n", maxLineNumWidth, d.LineNumber, d.LineOfCode)
		fmt.Fprintf(&builder, "|%s | %s%s\n", strings.Repeat(" ", maxLineNumWidth),
			strings.Repeat(" ", d.StartColumn),
			colorize(strings.Repeat("^", d.EndColumn-d.StartColumn), "green"))
		fmt.Fprintf(&builder, "|%s |\n", strings.Repeat(" ", maxLineNumWidth))
		fmt.Fprintf(&builder, "|%s \\ ---\n", strings.Repeat(" ", maxLineNumWidth))

		if i < len(db.diagnostics)-1 {
			fmt.Fprintf(&builder, "|\n")
		}
	}

	fmt.Fprintf(&builder, "\\ ---\n")

	return builder.String()
}
