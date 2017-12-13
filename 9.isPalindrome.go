package main

import (
	"fmt"
	"math"
)

//12521
//12521%10/1 = 1
//12521%100/10 = 2
//12521%1000/100 = 5
//12521%10000/1000 =2
//12521%100000/10000 = 1

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	f := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }
	y := x
	z := func(y int) int {
		z := 0
		for y > 0 {
			y = y / 10
			z++
		}
		return z
	}(y)

	for i := 0; i < z/2; i++ {
		if x%(f(10, (i+1)))/(f(10, i)) != x%(f(10, (z-i)))/f(10, (z-i-1)) {
			return false
		}
	}
	return true
}

func main() {
	x := 13575311
	fmt.Println("is a palindroem", x, isPalindrome(x))
}
