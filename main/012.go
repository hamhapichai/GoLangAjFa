package main

import "fmt"

func main() {
	var input int
	var temp int
	fmt.Print("Enter Integer : ")
	fmt.Scan(&input)
	fmt.Println(input)
	temp = input
	for a := 0; a < 4; a++ {
		input = input + 1
		temp = temp + input
		fmt.Println(temp)
	}
	fmt.Print("Result = ")
	fmt.Print(temp)
}
