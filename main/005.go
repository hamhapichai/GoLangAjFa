package main

import "fmt"

func main() {
	var first int
	var second int
	fmt.Print("Enter first integer : ")
	fmt.Scan(&first)
	fmt.Print("Enter second integer : ")
	fmt.Scan(&second)
	if first > second {
		fmt.Print("Largest int is ")
		fmt.Print(first)
	} else if second > first {
		fmt.Print("Largest int is ")
		fmt.Print(second)
	} else if first == second {
		fmt.Print("There are equal.")
	} else {
		fmt.Print("error")
	}
}
