package main

import (
	"fmt"
	"sort"
)

//time complexity is O(n^2 * logn)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	find := sort.SearchInts(nums, 0)
	fmt.Println(find, nums)
	length := len(nums)
	r := [][]int{}
	if len(nums) >= 3 && nums[0] == nums[length-1] && nums[0] == 0 {
		return [][]int{{0, 0, 0}}
	}
	if len(nums) < 3 || nums[length-1] < 0 || nums[0] > 0 || length == find {
		return r
	}
	for i := 0; i <= find; {
		for j := length - 1; j >= find; {
			s := nums[i] + nums[j]
			f := -1
			if i+1 <= j {
				fmt.Println("....", length, find, i+1, j)
				slc := nums[i+1 : j]
				f = sort.SearchInts(slc, -s)
				if f != len(slc) && slc[f] == -s {
					r = append(r, []int{nums[i], nums[j], -s})
				}
			}
			for j > 0 && nums[j] == nums[j-1] {
				j--
			}
			j--
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return r
}
func threeSumWrong(nums []int) [][]int {
	lcur, rcur := 0, len(nums)-1
	sort.Ints(nums)
	fmt.Println(nums)
	r := [][]int{}
	for lcur < rcur {
		fmt.Println(":::", lcur, rcur)
		sum := nums[lcur] + nums[rcur]
		find := sort.SearchInts(nums[lcur+1:rcur], -sum)
		if -sum == nums[lcur+1 : rcur][find] {
			f := nums[lcur+1 : rcur][find]
			r = append(r, []int{nums[lcur], nums[rcur], f})
			fmt.Println([]int{nums[lcur], nums[rcur], f})
		}
		if sum > 0 {
			tmp := lcur
			for lcur < rcur && nums[lcur] == nums[lcur+1] {
				lcur++
			}
			if tmp == lcur {
				lcur++
			}
		} else {
			tmp := rcur
			for lcur < rcur && nums[rcur] == nums[rcur-1] {
				rcur--
			}
			if tmp == rcur {
				rcur--
			}
		}
	}
	return r
}

func main() {
	//origin := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	//origin := []int{-1, 0, 1, 2, -1, -4}
	//origin := []int{}
	//origin := []int{0, 0, 0}
	//origin := []int{0}
	//origin := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	origin := []int{1, 1, -2}
	fmt.Println(origin)
	r := threeSum(origin)
	fmt.Println("good", r)
}
