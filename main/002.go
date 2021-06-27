package main

import "fmt"

func main() {
	var amount float32
	var price float32
	var tot float32
	var cash float32
	var change float32

	fmt.Print("Enter price of product : ")
	fmt.Scanln(&price)
	fmt.Print("Enter amount of product : ")
	fmt.Scanln(&amount)
	tot = amount * price
	fmt.Print("Total : ")
	fmt.Print(tot)
	fmt.Print(" Baht")
	fmt.Println("")
	fmt.Print("Cash recived : ")
	fmt.Scan(&cash)
	change = cash - tot
	fmt.Print("Change ")
	fmt.Print(change)
	fmt.Print(" Baht")
}
