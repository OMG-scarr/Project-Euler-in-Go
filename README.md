# Project Euler in Go

Solutions to Project Euler mathematical and computational problems, written in Go.

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)
![Problems Solved](https://img.shields.io/badge/Problems%20Solved-6-brightgreen?style=flat-square)

## About

A collection of solutions to [Project Euler](https://projecteuler.net/) problems implemented in Go. Each solution is a standalone program that can be compiled and run independently. Project Euler is a series of challenging mathematical/computer programming problems that require creative problem-solving beyond just mathematical insights.

## Solutions

| # | Problem | Approach | Solution |
|---|---------|----------|----------|
| 1 | [Multiples of 3 or 5](https://projecteuler.net/problem=1) | Iterates through all numbers below 1000, summing those divisible by 3 or 5 separately, then combining | [problem1.go](solutions/problem1.go) |
| 2 | [Even Fibonacci Numbers](https://projecteuler.net/problem=2) | Generates Fibonacci sequence up to 4,000,000 and sums only even-valued terms | [problem2.go](solutions/problem2.go) |
| 3 | [Largest Prime Factor](https://projecteuler.net/problem=3) | Trial division to extract all prime factors of 600,851,475,143, then identifies the largest | [problem3.go](solutions/problem3.go) |
| 4 | [Largest Palindrome Product](https://projecteuler.net/problem=4) | Brute-force multiplication of all 3-digit pairs, filters palindromes by string reversal, finds the maximum | [problem4.go](solutions/problem4.go) |
| 5 | [Smallest Multiple](https://projecteuler.net/problem=5) | Computes LCM(1..20) by tracking the maximum prime factor powers needed across all numbers | [problem5.go](solutions/problem5.go) |
| 6 | [Sum Square Difference](https://projecteuler.net/problem=6) | Computes the square of the sum and the sum of squares for 1..10, then takes the difference | [problem6.go](solutions/problem6.go) |

## Project Structure

```
Project-Euler-in-Go/
├── solutions/
│   ├── problem1.go    # Multiples of 3 or 5
│   ├── problem2.go    # Even Fibonacci Numbers
│   ├── problem3.go    # Largest Prime Factor
│   ├── problem4.go    # Largest Palindrome Product
│   ├── problem5.go    # Smallest Multiple (LCM)
│   └── problem6.go    # Sum Square Difference
├── LICENSE
└── README.md
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later (solutions use `range` over integers, introduced in Go 1.22)

## Running Solutions

Each solution is a standalone `package main` program. Run any solution with:

```bash
cd solutions
go run problem1.go
```

Or compile and run:

```bash
cd solutions
go build -o problem1 problem1.go
./problem1
```

To run all solutions at once:

```bash
cd solutions
for f in problem*.go; do echo "=== $f ===" && go run "$f"; done
```

## Contributing

To add a new solution:

1. Create a new file in `solutions/` named `problemN.go` (where N is the problem number)
2. Use `package main` with a `main()` function that prints the answer
3. Update the solutions table in this README

## License

MIT License - see [LICENSE](LICENSE) for details.
