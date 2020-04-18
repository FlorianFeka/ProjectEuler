package main

import "fmt"

var sum = 0
var target = 4_000_000

func main()  {
	// my simple approach
	fib(1,1)
	fmt.Println("SUM:",sum)
}

func fib(x, y int) int {
	fmt.Println("FIB:",y)
	if y % 2 == 0 {
		sum += y
	}
	x += y
	if(x > target) {
		return 0
	}
	return fib(y,x)
}