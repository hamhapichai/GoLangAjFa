package main

import "fmt"

/* GoLang doesn't have While Loop */

func main() {
	a := 1
	b := 0
	for {
		b = b + a
		a = a + 1
		if a >= 6 {
			goto NE
		}
	}
NE:
	fmt.Print("Result = ")
	fmt.Print(b)
}
