package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	x := 0
	str = strings.Trim(str, " \t")
	length := len(str)
	if length == 0 {
		return 0
	}
	sign := 1
	i := 0
	if str[0] == '+' {
		i = 1
	}
	if str[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < length; i++ {

		//fmt.Println("xxxx", string(str[i]), i)
		pos := int(str[i] - '0')
		if !(pos >= 0 && pos <= 9) {
			return x * sign
		}
		candid := (int64(x)*10 + int64(pos)) * int64(sign)
		if candid > 0 && candid >= int64(math.MaxInt32) {
			return math.MaxInt32
		} else if candid < 0 && candid <= int64(math.MinInt32) {
			return math.MinInt32
		}
		x = x*10 + pos
	}
	return x * sign
}

func main() {
	//i := myAtoi("999999999999999")
	//i := myAtoi("-2147483648")
	//i := myAtoi("2147483648") //expect: 2147483647
	i := myAtoi("-2147483649") //expect: -2147483648
	fmt.Println(math.MaxInt32, math.MinInt32)

	fmt.Println(i)
}
