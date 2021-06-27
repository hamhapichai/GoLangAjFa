package main

import "fmt"

func main() {
	var temp int
	sto := 0
	num := 0
	var stof float32
	var numf float32
	var avg float32
	for {
		fmt.Print("Enter integer (Enter 0 to end program) : ")
		fmt.Scan(&temp)
		sto = sto + temp
		num = num + 1
		if temp == 0 {
			fmt.Print("sum = ")
			fmt.Print(sto)
			stof = float32(sto)
			numf = float32(num)
			numf = numf - 1
			avg = stof / numf
			fmt.Print(". Average = ")
			fmt.Print(avg)
			break
		}
	}
}
