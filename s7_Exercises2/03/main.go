package main

import "fmt"

// create a TYPE constant and an UNTYPED constant

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Println(a, b)
}
