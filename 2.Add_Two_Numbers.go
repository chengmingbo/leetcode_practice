package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func whichLong(len1, len2 int) (int, int, int) {
	if len1 > len2 {
		return len1 - len2, len2, 1
	}
	return len2 - len1, len1, 2
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := length(l1), length(l2)
	diff, same, which := whichLong(len1, len2)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var lcur, scur *ListNode
	if which == 1 {
		lcur, scur = l1, l2
	} else {
		lcur, scur = l2, l1
	}
	carry := 0
	var s, head *ListNode
	for i := 0; i < same; i++ {
		//fmt.Println(lcur, scur)
		sum := lcur.Val + scur.Val + carry
		//fmt.Println(sum)
		carry = sum / 10
		res := sum % 10
		if s == nil {
			s = new(ListNode)
			head = s
			s.Next = nil
			s.Val = res
		} else {
			s.Next = new(ListNode)
			s.Next.Val = res
			s.Next.Next = nil
			s = s.Next
		}
		lcur, scur = lcur.Next, scur.Next
	}
	for i := 0; i < diff; i++ {
		sum := lcur.Val + carry
		carry = sum / 10
		res := sum % 10
		s.Next = new(ListNode)
		s.Next.Val = res
		s.Next.Next = nil
		s = s.Next
		lcur = lcur.Next
	}
	if carry != 0 {
		s.Next = new(ListNode)
		s.Next.Val = carry
		s.Next.Next = nil
	}
	fmt.Println("diff, length", diff, length(head))
	travel(head)
	return head

}

func add2Link(s *ListNode, arr []int) {
	if len(arr) <= 0 {
		return
	}
	s.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		s.Next = new(ListNode)
		s.Next.Val = arr[i]
		s.Next.Next = nil
		s = s.Next
	}
}

func length(s *ListNode) int {
	head := s
	i := 0
	for head != nil {
		i++
		head = head.Next
	}
	return i
}

func travel(s *ListNode) {
	head := s
	if head == nil {
		fmt.Println("null")
		return
	}
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	var s = new(ListNode)
	var t = new(ListNode)
	x := []int{1, 5, 2, 1, 3}
	y := []int{2, 1, 3}
	add2Link(s, x)
	add2Link(t, y)
	fmt.Println(length(s))
	fmt.Println(length(t))
	travel(s)
	travel(t)
	addTwoNumbers(s, t)

}
