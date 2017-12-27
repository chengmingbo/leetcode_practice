package main

import (
	"fmt"
)

/*
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

      After removing the second node from the end, the linked list becomes 1->2->3->5.
	  Note:
	  Given n will always be valid.
	  Try to do this in one pass.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func add2List(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	s := new(ListNode)
	p := s
	s.Val = a[0]
	s.Next = nil
	for _, v := range a[1:] {
		n := new(ListNode)
		n.Val = v
		n.Next = nil
		s.Next = n
		s = s.Next
	}
	return p
}

func travel(lst *ListNode) {
	if lst == nil {
		return
	}
	head := lst
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Printf("\b\b     \n")
	//fmt.Println()
}

//need travel one pass, use space substitute time, using an array save each node info, then compute remove which
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := []*ListNode{}
	lst := head
	i := 0
	for ; lst != nil; i++ {
		arr = append(arr, lst)
		lst = lst.Next
	}
	rm_idx := i - n
	fmt.Println("rm_idx, n, i", rm_idx, n, i)
	if rm_idx == 0 {
		return head.Next
	} else {
		arr[rm_idx-1].Next = arr[rm_idx].Next
	}
	return head
}
func main() {
	lst := []int{1, 3, 5, 6, 1}
	//lst := []int{1}
	var x *ListNode
	x = add2List(lst)
	fmt.Println("travel before")
	travel(x)
	y := removeNthFromEnd(x, 3)
	fmt.Println("travel after")
	travel(y)

}
