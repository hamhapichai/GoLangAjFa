package main

import "fmt"

func main() {
	var salary float32
	var workhour float32
	var tot float32
	var net float32
	fmt.Print("Enter salary rate (Baht/hour) : ")
	fmt.Scan(&salary)
	fmt.Print("Enter workhour (Hr.) : ")
	fmt.Scan(&workhour)
	tot = salary * workhour
	if tot > 5000 {
		net = tot - socialfee(salary)
	} else {
		net = tot
	}
	fmt.Print("Total = ", net)
}

func socialfee(num1 float32) float32 {
	var scf float32
	scf = num1 * 0.05
	return scf
}
