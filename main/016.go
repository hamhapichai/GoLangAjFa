package main

import "fmt"

func main() {
	var temp int
	sto := 0
	for {
		fmt.Print("Enter integer : ")
		fmt.Scan(&temp)
		sto = sto + temp
		if temp == 0 {
			fmt.Print("sum  = ")
			fmt.Print(sto)
			break
		}
	}
}
