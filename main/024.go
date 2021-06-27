package main

import "fmt"

func main() {
	var salary float32
	var workhour float32
	var tot float32
	var net float32
	fmt.Print("Enter salary rate (Baht/hour) : ")
	fmt.Scan(&salary)
	fmt.Print("Enter workhour (Hr./week) : ")
	fmt.Scan(&workhour)
	if workhour > 40 {
		var salait float32
		var salaot float32
		salait = salary * 40
		salaot = salary * (workhour - 40) * 1.15
		tot = salait + salaot
		net = tot - socialfee(tot)
		fmt.Print("Total salary = ", net)
	} else if workhour <= 40 {
		tot = salary * workhour
		net = tot - socialfee(tot)
		fmt.Print("Total salary = ", net)
	}
}

func socialfee(num1 float32) float32 {
	var scf float32
	scf = num1 * 0.05
	return scf
}
