package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var product int = 0
var palindrome_slice = []string{}

func palindrome_check(n int) {
	n_slice := []string{}

	n_string := strconv.Itoa(n)
	n_string_split := strings.Split(n_string, "")
	n_slice = append(n_slice, n_string_split...)
	n_sliceCopy := make([]string, len(n_slice))
	copy(n_sliceCopy, n_slice)
	slices.Reverse(n_sliceCopy)

	n_string_rev := strings.Join(n_sliceCopy, "")

	if n_string == n_string_rev {
		palindrome_slice = append(palindrome_slice, n_string)
	}
}

func largest_factors(palindrome string) {
	pal_num, _ := strconv.Atoi(palindrome)
	for i := 100; i < 1000; i++ {
		if pal_num%i == 0 {
			j := pal_num / i
			if j >= 100 && j < 1000 {
				fmt.Printf("Largest palindrome is %d = %d * %d\n", pal_num, i, j)
				return
			}
		}
	}
	fmt.Printf("No 2-digit factors found for palindrome %s\n", palindrome)
}

func palindrome_product() {
	palindrome_slice = []string{} // clear slice in case of multiple runs

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product = i * j
			palindrome_check(product)
		}
	}

	if len(palindrome_slice) == 0 {
		fmt.Println("No palindromes found")
		return
	}

	// Find the largest palindrome (as number)
	max_palindrome := 0
	for _, s := range palindrome_slice {
		n, _ := strconv.Atoi(s)
		if n > max_palindrome {
			max_palindrome = n
		}
	}
	largest_palindrome := strconv.Itoa(max_palindrome)
	largest_factors(largest_palindrome)
}

func main() {
	palindrome_product()
}
