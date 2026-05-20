package main

import (
	"fmt"
)

func specialPythagoreanTriplet(n int) {
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			c := float64(n) - float64(a) - float64(b)
			if c > 0 && float64(a*a)+float64(b*b) == c*c {
				fmt.Println(a, b, c, a*b*int(c))
			}
		}
	}
}

func main() {
	specialPythagoreanTriplet(1000)
}
