package main

import "fmt"

func main() {
	var amount float32
	var price float32
	var tot float32

	fmt.Print("Enter price of product : ")
	fmt.Scanln(&price)
	fmt.Print("Enter amount of product : ")
	fmt.Scanln(&amount)
	tot = amount * price
	fmt.Print("Total : ")
	fmt.Print(tot)
	fmt.Print(" Baht")
}
