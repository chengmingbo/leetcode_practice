package main

import (
	"fmt"
)

type SumEle struct {
	Idx1 int
	Idx2 int
}

type Trip struct {
	Ele1 int
	Ele2 int
	Ele3 int
}

type MS map[int][]SumEle

func (se SumEle) isNotin(i int) bool {
	return i != se.Idx1 && i != se.Idx2
}

func (t *Trip) sortT() {
	if t.Ele1 > t.Ele2 {
		t.Ele1, t.Ele2 = t.Ele2, t.Ele1
	}
	if t.Ele2 > t.Ele3 {
		t.Ele2, t.Ele3 = t.Ele3, t.Ele2
	}
	if t.Ele1 > t.Ele3 {
		t.Ele1, t.Ele3 = t.Ele3, t.Ele1
	}
	if t.Ele1 > t.Ele2 {
		t.Ele1, t.Ele2 = t.Ele2, t.Ele1
	}
}

func threeSum(nums []int) [][]int {

	//flatten nums two sum dimension to n(n-1)/2 one.
	m := make(MS)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			s := nums[i] + nums[j]
			ele := SumEle{i, j}
			m[s] = append(m[s], ele)
		}
	}
	s := make(map[Trip]int)
	//find a solution
	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		neg := -nums[i]
		if vs, ok := m[neg]; ok {
			for j := 0; j < len(vs); j++ {
				v := vs[j]
				if !v.isNotin(i) {
					continue
				}
				t := Trip{nums[v.Idx1], nums[v.Idx2], nums[i]}
				t.sortT()
				if _, ok := s[t]; ok {
					//fmt.Println(t)
					continue
				}
				s[t] = 1
				//fmt.Println(t)
				outcome := []int{nums[v.Idx1], nums[v.Idx2], nums[i]}
				ret = append(ret, outcome)
			}
		}
	}

	return ret
}

func main() {
	//r := threeSum([]int{-1, 0, 1, 2, -1, -4})
	oringin := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(oringin)
	r := threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6})
	fmt.Println("good", r)
}
