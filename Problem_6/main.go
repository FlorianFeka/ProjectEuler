package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumSquareDifference(100))
}

func sumSquareDifference(target float64) int {
	squareSum := (target * (target + 1)) / 2
	squareSum = math.Pow(squareSum, 2)
	sumSquare := (target * (target + 1) * (2 * target + 1)) / 6

	return int(squareSum - sumSquare)
}
