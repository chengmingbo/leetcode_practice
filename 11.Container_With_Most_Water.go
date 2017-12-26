package main

import (
	"fmt"
	"math"
)

//a O(n^2) solution
func maxAreaSlow(height []int) int {
	maxA := 0
	minInt := func(a, b int) int { return int(math.Min(float64(a), float64(b))) }
	for idx, h1 := range height {
		for jdx, h2 := range height {
			if idx <= jdx {
				continue
			}
			area := (idx - jdx) * minInt(h1, h2)
			if area > maxA {
				maxA = area
			}
		}
	}
	return maxA
}

//two pointer, find a shorter one, then move a step
func maxArea(height []int) int {
	length := len(height)
	bp, ep := 0, length-1
	area := -1
	minInt := func(a, b int) int { return int(math.Min(float64(a), float64(b))) }
	for bp < ep {
		wide := ep - bp
		h1, h2 := height[bp], height[ep]
		h := minInt(h1, h2)
		na := wide * h
		if na > area {
			area = na
		}
		if h1 < h2 {
			bp++
		} else {
			ep--
		}
	}
	return area
}
func main() {
	lst := []int{1, 2, 1}
	lst := []int{1, 2, 4, 3} //expected 4
	a := maxArea(lst)
	fmt.Println("ok.")
	fmt.Println(a)
}
