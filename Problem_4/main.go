package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Started ...")
	fmt.Println(largestPalindromeProduct(100, 999))
}

func largestPalindromeProduct(from, to int) int {
	largestPalindrome := 0

	for i := from; i <= to; i++ {
		for j := from; j <= to; j++ {
			product := i * j
			if isPalindrome(strconv.Itoa(product)) && product > largestPalindrome {
				largestPalindrome = product
			}
		}
	}

	return largestPalindrome
}

func isPalindrome(number string) bool {
	for i := 0; i < len(number); i++ {
		if number[i] != number[len(number)-1-i] {
			return false
		}
	}

	return true
}
