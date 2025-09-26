package main

import (
	"fmt"
	"sort"
)

func prime_factor(n int) (int, int) {
	primes := []int{}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
		}
	}
	sort.SliceStable(primes, func(i, j int) bool {
		return true
	})
	fmt.Println(primes)
	return primes[0], primes[1]
}

// achieved but there are exceptions not handled
func main() {
	fmt.Println(prime_factor(600851475143))
	// fmt.Println(prime_factor(1120))
}
