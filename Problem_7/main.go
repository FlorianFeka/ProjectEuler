package main

import "fmt"

func main() {
	// my simple soultion
	count := 0
	currentPrime := 0
	for i := 2; count < 10_001; i++ {
		prime := true
		for j := 2; j <= i; j++ {
			if i != j && i%j == 0{
				prime = false
			}
			
		}
		if prime {
			count ++
			currentPrime = i
		}
	}
	fmt.Println(currentPrime)
}