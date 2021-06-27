package main

import "fmt"

func main() {
	var poppu int
	var year int
	var fpoppu float32
	fmt.Print("Popualtion : ")
	fmt.Scan(&poppu)
	fmt.Print("How many year you want to calculate : ")
	fmt.Scan(&year)
	fpoppu = float32(poppu)
	for a := 0; a < year; a++ {
		fpoppu = fpoppu * 1.1
	}
	fmt.Print("Total population = ", fpoppu)
}
