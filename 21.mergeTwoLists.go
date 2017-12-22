package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildAList(li []int) *ListNode {
	if len(li) == 0 {
		return nil
	}
	ln := new(ListNode)
	head := ln
	ln.Next = nil
	ln.Val = li[0]
	for _, e := range li[1:] {
		nln := new(ListNode)
		nln.Val = e
		nln.Next = nil
		ln.Next = nln
		ln = ln.Next
	}
	return head
}

func travel(lin *ListNode) {
	if lin == nil {
		return
	}
	head := lin
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Printf("\b\b   \n")
}

func length(lin *ListNode) int {
	if lin == nil {
		return 0
	}
	head := lin
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := new(ListNode)
	head := ln

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head.Val = l1.Val
		l1 = l1.Next
	} else {
		head.Val = l2.Val
		l2 = l2.Next
	}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			//clean l2 and break
			for l2 != nil {
				lin := new(ListNode)
				lin.Val = l2.Val
				lin.Next = nil
				head.Next = lin
				head = head.Next
				l2 = l2.Next
			}
			return ln

		} else if l2 == nil {
			//clean l1 and break
			for l1 != nil {
				lin := new(ListNode)
				lin.Val = l1.Val
				lin.Next = nil
				head.Next = lin
				head = head.Next
				l1 = l1.Next
			}
			return ln
		}

		//l1 & l2 not empty: add the bigger
		lin := new(ListNode)
		if l1.Val > l2.Val {
			lin.Val = l2.Val
			l2 = l2.Next
		} else {
			lin.Val = l1.Val
			l1 = l1.Next
		}
		lin.Next = nil
		head.Next = lin
		head = head.Next
		//if only l1 or only l2, add
	}

	return ln
}

func main() {
	fmt.Println("")
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}

	lin1 := buildAList(l1)
	lin2 := buildAList(l2)
	//travel(lin)
	lin := mergeTwoLists(lin1, lin2)
	travel(lin)
	n := length(lin)
	fmt.Println("length:", n)
}
