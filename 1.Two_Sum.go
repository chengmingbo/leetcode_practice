package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, x := range nums {
		m[x] = i
	}
	fmt.Printf("%v\n", m)

	for j, x := range nums {
		diff := target - x
		println(j, diff)
		if v, ok := m[diff]; ok && j != v {
			return []int{j, v}
			fmt.Printf("%v, %v\n", j, v)
		}
	}

	return []int{0, 0}

}
func main() {
	a := []int{3, 2, 4}
	b := 6
	twoSum(a, b)
	fmt.Println(twoSum(a, b))
}
