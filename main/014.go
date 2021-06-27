package main

import "fmt"

/* GoLang doesn't have while loop */

func main() {
	var input int
	var temp int
	var loop int
	a := 0
	fmt.Print("Enter Integer : ")
	fmt.Scan(&input)
	fmt.Print("Enter round of loop you want : ")
	fmt.Scan(&loop)
	loop = loop - 1
	fmt.Println(input)
	temp = input
	for a < loop {
		input = input + 1
		temp = temp + input
		fmt.Println(temp)
		a = a + 1
	}
	fmt.Print("Result = ")
	fmt.Print(temp)
}
