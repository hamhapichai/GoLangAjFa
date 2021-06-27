package main

import "fmt"

func main() {
	var high float32
	var width float32
	var size float32
	fmt.Print("Enter triangle high : ")
	fmt.Scan(&high)
	fmt.Print("Enter triangle width : ")
	fmt.Scan(&width)
	size = 0.5 * high * width
	fmt.Print("Size of triangle = ")
	fmt.Print(size)
}
