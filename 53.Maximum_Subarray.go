package main

import (
	"fmt"
	"math"
)

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

========[-2, 1, -3, 4, -5, 2, 1, -1, 4]======num[i],  sumi,  maxi
curr1:       |   ^                              1       1     1
curr2:           |  ^                          -3      -2     1
curr3:              |   ^                       4       4     4
curr4:                  | ^                    -5      -1     4
curr5:                    |   ^                 2       2     4
curr6:                        |  ^              1       3     4
curr7:                           |  ^          -1       2     4
curr8:                              |           4       6     6
*/

//from https://leetcode.com/problems/maximum-subarray/discuss/20193
//sum(0,i) = a[i] + (sum(0,i-1) < 0 ? 0 : sum(0,i-1))

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxInt := func(a, b int) int { return int(math.Max(float64(a), float64(b))) }
	maxi := nums[0]
	sumi := nums[0]
	fmt.Println("list\t\t\t\tnums[i], sumi, maxi")
	for i := 1; i < length; i++ {
		sumi = maxInt(nums[i], sumi+nums[i])
		maxi = maxInt(sumi, maxi)
		fmt.Printf("%d:\t", nums)
		fmt.Println(nums[i], sumi, maxi)
	}
	return maxi
}

func main() {
	//x := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	x := []int{-2, 1, -3, 4, -5, 2, 1, -1, 4}
	y := maxSubArray(x)
	fmt.Println("xxxx: ", y)
}
