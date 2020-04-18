package main

import "fmt"

var sum = 0
var target = 4_000_000

func main()  {
	// my simple approach
	fib(1,1)
	fmt.Println(sum)
}

func fib(x, y int) int {
	if(y > target){
		return x+y
	}
	fmt.Println(y)
	if y % 2 == 0 {
		sum += y
	}
	x += y
	return fib(y,x)
}