package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func palindrome() {
	n := 0
	n_slice := []string{}
	n_string := strconv.Itoa(n)
	fmt.Println(n_string)
	n_string_split := strings.Split(n_string, "")
	n_slice = append(n_slice, n_string_split...)
	n_sliceCopy := make([]string, len(n_slice))
	copy(n_sliceCopy, n_slice)
	slices.Reverse(n_sliceCopy)
	fmt.Println(n_sliceCopy)
	n_string_rev := strings.Join(n_sliceCopy, "")

	if n_string == n_string_rev {
		fmt.Println("This is a palindrome!")
	} else {
		fmt.Println("This is nowhere close to a palindrome")
	}

}

func main() {
	palindrome()
}
