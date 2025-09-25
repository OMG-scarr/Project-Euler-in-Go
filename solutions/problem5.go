//LCM (Least Common Multiple)

package main

import "fmt"

// Returns a slice of prime factors for n
func prime_factors(n int) []int {
	var factors []int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

// Helper: counts the occurrences of each prime in a slice
func count_factors(factors []int) map[int]int {
	counts := make(map[int]int)
	for _, f := range factors {
		counts[f]++
	}
	return counts
}

// Computes the LCM from the prime factors of all numbers 2..n
func lcm_from_factors(n int) int {
	max_counts := make(map[int]int)
	for i := 2; i <= n; i++ {
		counts := count_factors(prime_factors(i))
		for prime, count := range counts {
			if count > max_counts[prime] {
				max_counts[prime] = count
			}
		}
	}
	lcm := 1
	for prime, count := range max_counts {
		for i := 0; i < count; i++ {
			lcm *= prime
		}
	}
	return lcm
}

func main() {
	fmt.Println(lcm_from_factors(20))
}
