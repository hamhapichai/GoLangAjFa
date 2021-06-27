package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter int : ")
	fmt.Scan(&input)
	for a := 0; a < input; a++ {
		for b := 1; b < (a + 2); b++ {
			fmt.Print("# ")
		}
		fmt.Print("\n")
	}
}
