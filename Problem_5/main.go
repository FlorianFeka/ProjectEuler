package main

import "fmt"

func main() {
	fmt.Println(smallestMultiple(20))
}

func smallestMultiple(target int) int {
	for i := 2; ; i++ {
		if isDividable(i, target) {
			return i
		}
	}
}

func isDividable(number, target int) bool {
	for i := 1; i < target; i++ {
		if number%i != 0 {
			return false
		}
	}

	return true
}
