# Einführung in **Kvexium**

**Kvexium** ist eine moderne Programmiersprache, die für Entwickler entwickelt wurde, die klare, präzise und leistungsstarke Programme schreiben möchten. Die Sprache kombiniert eine einfache, lesbare Syntax mit fortschrittlichen Funktionen, die es dir ermöglichen, komplexe Aufgaben effizient zu lösen.

### DOKUMENTATION

### Warum **Kvexium**?

- **Klar und prägnant**: Die Syntax von Kvexium ist intuitiv und einfach zu verstehen.
- **Moderne Kontrollstrukturen**: Mächtige `if`, `for`, `while` und `foreach`-Schleifen erleichtern das Schreiben von sauberem Code.
- **Präzise Datentypen**: Unterstützt erweiterte Typen für hochpräzise Berechnungen.

### Beispielcode

```kvexium
#Main {
	dec x: i8;
	x := 9;
	x--;

	if (x > 5) {
		print("'x' ({x}) is bigger than 5");
	} : {
		print("'x' ({x}) isn't bigger than 5");
	}

	for (i: i8 = 0; i <= 8; i++) {
		print("i ({i}) is smaller or equal to 8\n");
	}

	dec b[]: i16[x] := {1, 5};

	for (i: i8 :: b) {
		print("item ({i}) is in the array named 'b'\n");
	}
}
```

### Zusammenfassung

**Kvexium** macht das Programmieren einfach und leistungsstark. Probiere es aus und erlebe, wie effizient du komplexe Aufgaben mit einer klaren und modernen Syntax lösen kannst!