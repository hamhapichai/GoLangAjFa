package main

import "fmt"

func main() {
	var temp int
	even := 0
	odd := 0
	for {
		fmt.Print("Enter integer : ")
		fmt.Scan(&temp)
		if (temp%2) == 0 && temp != 0 {
			even = even + temp
		} else if (temp % 2) == 1 {
			odd = odd + temp
		} else if temp == 0 {
			fmt.Print("sum of even = ")
			fmt.Print(even)
			fmt.Print(", sum of odd = ")
			fmt.Print(odd)
			break
		}
	}
}
