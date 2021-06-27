package main

import "fmt"

func main() {
	var salary float32
	var workhour float32
	var tot float32
	fmt.Print("Enter salary rate (Baht/hour) : ")
	fmt.Scan(&salary)
	fmt.Print("Enter workhour (Hr.) : ")
	fmt.Scan(&workhour)
	tot = salary * workhour
	fmt.Print("Total = ", tot)
}
