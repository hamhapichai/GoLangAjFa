package main

import "fmt"

/* GoLang doesn't have while loop */

func main() {
	var input int
	var temp int
	a := 0
	fmt.Print("Enter Integer : ")
	fmt.Scan(&input)
	fmt.Println(input)
	temp = input
	for {
		input = input + 1
		temp = temp + input
		fmt.Println(temp)
		a = a + 1
		if a >= 4 {
			goto NE
		}
	}
NE:
	fmt.Print("Result = ")
	fmt.Print(temp)
}
