package main

import (
	"fmt"
)

const n = 1000

var sum3 int = 0
var sum5 int = 0

/*
	func multiples3() {
		for i := range n {
			if (i % 3) == 0 {
				sum3 += i
			}
		}
		fmt.Println(sum3)
	}

	func multiples5() {
		sum5 := 0
		for i := range n {
			if (i % 5) == 0 {
				sum5 += i
			}
		}
		fmt.Println(sum5)
	}
*/
func main() {
	fmt.Println("Welcome to the ultimate test")
	for i := range n {
		if (i % 3) == 0 {
			sum3 += i
		}
	}
	fmt.Println(sum3)

	for i := range n {
		if (i % 5) == 0 {
			sum5 += i
		}
	}
	fmt.Println(sum5)

	var result int = (sum3 + sum5)

	fmt.Println(result)
}
