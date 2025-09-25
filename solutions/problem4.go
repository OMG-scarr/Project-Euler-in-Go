package main

import (
	"fmt"
	"strconv"
	"strings"
)

func palindrome(n int) bool {
	n_string := strconv.Itoa(n)
	n_string_split := strings.Split(n_string, "")
	n_sliceCopy := make([]string, len(n_string_split))
	copy(n_sliceCopy, n_string_split)
	for i, j := 0, len(n_sliceCopy)-1; i < j; i, j = i+1, j-1 {
		n_sliceCopy[i], n_sliceCopy[j] = n_sliceCopy[j], n_sliceCopy[i]
	}
	n_string_rev := strings.Join(n_sliceCopy, "")
	if n_string == n_string_rev {
		return true
	}
	return false
}

func main() {
	
	palindrome_Product := 0
	a, b := 0, 0

	// Ascending order loops for factors
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			product := i * j
			if palindrome(product) && product > palindromeMax {
				palindrome_Product = product
				a = i
				b = j
			}
		}
	}
	fmt.Printf("Largest palindrome is %d = %d * %d\n", palindromeMax, a, b)
}
