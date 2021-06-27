package main

import "fmt"

func main() {
	b := 0
	for a := 0; a < 6; a++ {
		b = b + a
	}
	fmt.Print("Result = ")
	fmt.Print(b)
}
