package main

import "fmt"

func main() {
	var fir int     /* เลขตัวแรก */
	var sec int     /* เลขตัวหลัง*/
	var store []int /* slice เก็บชุดตัวเลข */
	var diff int    /* ผลต่าง เอามาทำลูป */
	var temp int    /* ตัวแปรเก็บค่าชั่วคราวของ loop 1 */
	plus := 1
	fmt.Print("Enter first int : ")
	fmt.Scan(&fir)
	fmt.Print("Enter second int : ")
	fmt.Scan(&sec)
	store = append(store, fir)
	diff = sec - fir
	for a := 0; a < diff; a++ {
		temp = fir + plus
		store = append(store, temp)
		plus = plus + 1
	}
	fmt.Print("data = ", store)
	temp2 := 0                /* ตัวแปรพักค่าของ loop2 */
	for _, v := range store { /* loop ตามจำนวนสมาชิก โดยใช้ v ในการพักค่าของสมาชิกในแต่ละรอบ */
		if v%2 == 1 {
			temp2 = temp2 + v
		}
	}

	temp3 := 0
	for _, v := range store {
		if v%2 == 0 {
			temp3 = temp3 + v
		}
	}

	var sqstore []int
	var temp4 int
	for _, v := range store {
		temp4 = v * v
		sqstore = append(sqstore, temp4)
	}

	temp5 := 0
	for _, v := range sqstore {
		temp5 = temp5 + v
	}

	var choice string
	fmt.Print("\nMake a choice")
	fmt.Print("\nPress A : Sum of odd number.")
	fmt.Print("\nPress B : Sum of even number.")
	fmt.Print("\nPress C : Square of number.")
	fmt.Print("\nPress D : Sum of square.")
	fmt.Print("\n")
	fmt.Scan(&choice)
	switch choice {
	case "a":
		fmt.Print("result = ", temp2)
	case "b":
		fmt.Print("B. result = ", temp3)
	case "c":
		fmt.Print("C. result = ", sqstore)
	case "d":
		fmt.Print("D. result = ", temp5)
	default:
		fmt.Print("Invalid Choice.")
	}
}
