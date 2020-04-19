package main

import "fmt"

func main() {
	// my simple approach
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(target int) int {
	factor := 0
	for i := 2; i <= target; i++ {
		if target%i == 0 {
			target /= i
			if factor < i {
				factor = i
			}
		}
	}
	return factor
}
