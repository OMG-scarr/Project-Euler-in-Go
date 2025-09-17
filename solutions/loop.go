package main

import (
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

}

func palindrome_product() {

	for i := 10; i < 100; i++ {
		for j := 10; j < 100; j++ {
			product = i * j
			palindrome_check(product)

		}

	}

	largest_palindrome := palindrome_slice[len(palindrome_slice)-1]
	largest_factors(largest_palindrome)

}

func main() {
	palindrome_product()
}
