package main

import (
	"fmt"
	"math"
)

func reverse(xx int) int {

	//catch overflow condition
	var x int32 = int32(xx)
	if int(xx) != int(x) {
		return 0
	}

	var y int32 = 0
	var flag int32 = 1
	//ignore sign when reverse
	if x < 0 {
		flag = -1
		x = -x
	}

	for x > 0 {
		//if overflow, return 0
		if (int(y)*10 + int(x%10)) > int(math.MaxInt32) {
			return 0
		}

		//reverse
		y = y*10 + x%10
		x = x / 10
	}

	return int(y * flag)
}

func main() {
	ret := reverse(1534236469)
	fmt.Printf("%T %d\n", ret, ret)
}
