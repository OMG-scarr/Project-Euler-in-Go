package main

import "fmt"

func nthPrimeNumber(x int) int {
	primes := []int{}
	n := 2

	for len(primes) < x {
		isPrime := true

		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n)
		}
		n++
	}
	return primes[x-1]
}

func main() {
	fmt.Println(nthPrimeNumber(10))
}
