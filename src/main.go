package main

func main () {
	bytes, _ := os.readFile("./examples/00.kvxm")
	source := string(bytes)

	fmt.Printf("Code: {%s}\n", source)
}