package main

import (
	"fmt"
)

/*
Write a program that prints a number in decimal, bianary, and hex
*/

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
