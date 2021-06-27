package main

import "fmt"

func main() {
	var score float32
	fmt.Print("Enter your score : ")
	fmt.Scan(&score)
	if score >= 80 {
		fmt.Print("Grade : A")
	} else if (score < 80) && (score >= 75) {
		fmt.Print("Grade : B+")
	} else if (score < 75) && (score >= 70) {
		fmt.Print("Grade : B")
	} else if (score < 70) && (score >= 65) {
		fmt.Print("Grade : C+")
	} else if (score < 65) && (score >= 60) {
		fmt.Print("Grade : C")
	} else if (score < 60) && (score >= 55) {
		fmt.Print("Grade : D+")
	} else if (score < 55) && (score >= 50) {
		fmt.Print("Grade : D")
	} else if score < 50 {
		fmt.Print("Grade : F")
	} else {
		fmt.Print("error")
	}
}
