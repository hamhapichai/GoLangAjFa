package main

import "fmt"

func main() {
	var poppu int
	var year int
	var afbpoppu float32
	fmt.Print("Popualtion : ")
	fmt.Scan(&poppu)
	fmt.Print("How many year you want to calculate : ")
	fmt.Scan(&year)
	afbpoppu = born(poppu, year)
	totpop := die(afbpoppu, year)
	fmt.Print("Total population = ", totpop)
}

func born(popp int, y int) float32 {
	var fpoppu float32
	fpoppu = float32(popp)
	for a := 0; a < y; a++ {
		fpoppu = fpoppu * 1.1
	}
	return fpoppu
}

func die(peo float32, ye int) float32 {
	var dpoppu float32
	for b := 0; b < ye; b++ {
		dpoppu = float32(peo) * 0.04
		peo = peo - dpoppu
	}
	return peo
}
