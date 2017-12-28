package main

import (
	"container/list"
	"fmt"
)

/*
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.


Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//we can use a queue, add all node to the queue, including those nil node
//then compare each elements of two queues
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	fmt.Println("xxx")
	que := list.New()
	que.PushBack(p)
	lst1 := list.New()
	for que.Len() != 0 {
		node := que.Front().Value.(*TreeNode)
		if node != nil {
			que.PushBack(node.Left)
			que.PushBack(node.Right)
		}
		que.Remove(que.Front())
		lst1.PushBack(node)
	}

	lst2 := list.New()
	que = list.New()
	que.PushBack(q)

	for que.Len() != 0 {
		node := que.Front().Value.(*TreeNode)
		if node != nil {
			que.PushBack(node.Left)
			que.PushBack(node.Right)
		}
		que.Remove(que.Front())
		lst2.PushBack(node)
	}

	for h1, h2 := lst1.Front(), lst2.Front(); h1 != nil && h2 != nil; h1, h2 = h1.Next(), h2.Next() {
		if h1.Value.(*TreeNode) != nil && h2.Value.(*TreeNode) != nil {
			fmt.Println("h1.Val, h2.Val", h1.Value.(*TreeNode).Val, h2.Value.(*TreeNode).Val)
			if h1.Value.(*TreeNode).Val != h2.Value.(*TreeNode).Val {
				return false
			}
		} else if h1.Value.(*TreeNode) == nil && h2.Value.(*TreeNode) != nil {
			return false
		} else if h1.Value.(*TreeNode) != nil && h2.Value.(*TreeNode) == nil {
			return false
		}
	}
	return true
}

func genTreeNode(a int) *TreeNode {
	n := new(TreeNode)
	n.Val = a
	n.Left = nil
	n.Right = nil
	return n
}
func buildaTree(a []int) *TreeNode {
	t := new(TreeNode)
	t.Val = a[0]
	t.Left = genTreeNode(a[1])
	t.Right = genTreeNode(a[2])
	return t
}

func main() {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	//y := []int{1, 3, 2}
	a := buildaTree(x)
	b := buildaTree(y)
	fmt.Println(isSameTree(a, b))
	return
}
