package main

import "fmt"

var target = 999

func main() {
	// my simple solution
	sum := 0
	for i := 1; i < 1_000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// Euler Project solution (https://projecteuler.net/overview=001)
	res := SumDivisibleBy(3) + SumDivisibleBy(5) - SumDivisibleBy(15)
	fmt.Println(res)
}

func SumDivisibleBy(n int) int {
	p := target / n
	return n * (p * (p + 1)) / 2
}
