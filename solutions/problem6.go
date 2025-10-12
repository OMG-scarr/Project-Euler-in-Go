package main

import "fmt"

func sumSquareDifference(n int) {
	sum := 0
	sumSq := 0
	sqSum := 0

	for i := 1; i <= 10; i++ {
		sum += i
		sqSum += (i * i)
	}
	sumSq = sum * sum
	difference := sumSq - sqSum
	fmt.Println("Square of sum: ", sumSq, "Sum of Squares: ", sqSum, "Difference: ", difference)
}

func main() {
	sumSquareDifference(10)
}
