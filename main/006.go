package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter Int : ")
	fmt.Scan(&input)
	if (input % 2) == 0 {
		fmt.Print("This int is even number.")
	} else if (input % 2) == 1 {
		fmt.Print("This int is odd number.")
	} else {
		fmt.Print("error")
	}
}
