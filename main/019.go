package main

import "fmt"

func main() {
	var store []int
	var temp int
	for {
		fmt.Print("Enter Int : ")
		fmt.Scan(&temp)
		if temp != 0 {
			store = append(store, temp)
		} else if temp == 0 {
			min := store[0]
			for _, v := range store {
				if min < v {
					min = v
				}
			}
			max := store[0]
			for _, v := range store {
				if max > v {
					max = v
				}
			}
			fmt.Print("The largest int is ", max, ", The smallest int is ", min)
			break
		}
	}
}
