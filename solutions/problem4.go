package main

import (
	"slices"
	"strconv"
	"strings"
)

/*
	func palindrome_factors() (first, second int) {
		factors := []int{}
		two_digit_factors := []int{}
		for i := 1; i <= palindrome; i++ {
			if palindrome%i == 0 {
				factors = append(factors, i)
				if i < 100 {
					two_digit_factors = append(two_digit_factors, i)
				}
			}

		}
		sort.SliceStable(two_digit_factors, func(i, j int) bool { return true })
		fmt.Println(two_digit_factors[0], two_digit_factors[1])
		return two_digit_factors[0], two_digit_factors[1]
	}
*/
func palindrome_product() {
	first, second := 20, 13
	product := first * second
	palindrome_string := strconv.Itoa(product)
	split_palindrome := strings.Split(palindrome_string, "")
	slices.Reverse(split_palindrome)

}

func main() {
	palindrome_product()
}
