package main

import "fmt"

func main() {
	var temp int
	for {
		fmt.Print("Enter integer : ")
		fmt.Scan(&temp)
		if temp == 0 {
			fmt.Print("Input integer = 0, program ended")
			break
		}
	}
}
