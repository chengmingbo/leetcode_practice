package main

import (
	"fmt"
)

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"

]
*/

//Suppose we have many pairs of brackets, we can insert each of the bracket in to each position of a existed string, but how to delete replications
// if using travel all possible, it is a NP Hard Problem
func generateParenthesis(n int) []string {
	m := make(map[string]bool)         // record all right brackets
	amax_len := make(map[int][]string) //record the longest bracket
	if n < 1 {
		return []string{}
	}
	m["()"] = true
	amax_len[1] = []string{"()"}
	for i := 1; i < n; i++ {
		for _, v := range amax_len[i] {
			leng := len(v)
			fmt.Println("length", leng)
			for j := 0; j < leng; j++ {
				s := v[0:j] + "()" + v[j:]
				//fmt.Println(s)
				_, ok := m[s]
				if !ok {
					m[s] = true
					amax_len[i+1] = append(amax_len[i+1], s)
				}
			}
		}
	}
	return amax_len[n]
}

func main() {
	x := generateParenthesis(5)
	fmt.Println("Hello World", x)
}
