package main

import "fmt"

func main() {
	var poppu int
	var fpoppu float32
	fmt.Print("Popualtion : ")
	fmt.Scan(&poppu)
	fpoppu = float32(poppu)
	count := 0
	for {
		if fpoppu >= 100000 {
			count = 0
			break
		}
		fpoppu = fpoppu * 1.1
		count = count + 1
		if fpoppu > 100000 {
			break
		}
	}
	fmt.Print("Time to 100k = ", count, " year(s).")
}
