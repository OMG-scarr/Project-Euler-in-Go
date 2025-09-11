package main

import "fmt"

func main() {

	a := 1
	b := 2
	sum := 0

	for b < 4000000 {
		if b%2 == 0 {
			sum += b
		}
		temp := a + b
		a = b
		b = temp
	}

	fmt.Println(sum)
}
