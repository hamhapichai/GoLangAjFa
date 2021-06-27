package main

import "fmt"

func main() {
	var high float32
	var width float32
	var size float32
	fmt.Print("Enter rectangle high : ")
	fmt.Scan(&high)
	fmt.Print("Enter retangle width : ")
	fmt.Scan(&width)
	size = high * width
	fmt.Print("Retangle size = ")
	fmt.Print(size)
}
