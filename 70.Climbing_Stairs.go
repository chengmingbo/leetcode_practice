package main

import (
	"fmt"
)

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.


Example 1:

Input: 2
Output:  2
Explanation:  There are two ways to climb to the top.

1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output:  3
Explanation:  There are three ways to climb to the top.

1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	//add 2 if the last step is 1
	//add 1 if the last step is 2
	nums1 := 1
	nums2 := 0
	sum := 1
	for i := 1; i < n; i++ {
		sum = nums1*2 + nums2
		old2 := nums2
		nums2 = nums1
		nums1 = nums1 + old2
	}
	return sum
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, climbStairs(i))
	}
}
