package main

import (
	"fmt"
	//	"strconv"
)

var r2a = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	length := len(s)
	for i := 0; i < length; i++ {
		if i+1 < length && r2a[string(s[i+1])] > r2a[string(s[i])] {
			sum -= r2a[string(s[i])]
		} else {
			sum += r2a[string(s[i])]
		}
	}
	return sum
}

func main() {
	fmt.Println(r2a)
	fmt.Println(romanToInt("XCIX"))
}
