package main

import (
	"fmt"
)

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

type Stack struct {
	Item string
	Next *Stack
	Pre  *Stack
}

func pop(s *Stack) (*Stack, string) {
	if s == nil {
		return nil, ""
	}
	r := s.Item
	s = s.Pre
	if s == nil { //if s is nil, we can not find .Next anymore
		return s, r
	}
	s.Next.Pre = nil
	s.Next = nil
	return s, r
}

func push(s *Stack, v string) *Stack {
	if s == nil {
		p := new(Stack)
		p.Item = v
		p.Next = nil
		p.Pre = nil
		s = p

		return s
	}
	p := new(Stack)
	p.Item = v
	p.Next = nil
	p.Pre = s

	s.Next = p
	s = s.Next

	return s
}

func isEmpty(s *Stack) bool {
	if s == nil {
		return true
	}
	return false
}

func top(s *Stack) string {
	if s == nil {
		return ""
	}
	return s.Item
}

func length(s *Stack) int {
	n := 0
	for s != nil {
		n++
		s = s.Pre
	}
	return n
}

// use a stack, to decide wether a bracket finishes reasonable, if not return false, else, throw these
func isValid(s string) bool {
	sk := new(Stack)
	sk = nil
	for _, re := range s {
		e := string(re)
		fmt.Println("new:", e)
		var v string
		if e == ")" {
			if sk, v = pop(sk); v != "(" {
				return false
			}
			continue
		} else if e == "}" {
			if sk, v = pop(sk); v != "{" {
				return false
			}
			continue
		} else if e == "]" {
			if sk, v = pop(sk); v != "[" {
				return false
			}
			continue
		}
		sk = push(sk, string(e))
		fmt.Println("push", e)
	}
	if isEmpty(sk) {
		return true
	}
	return false

}

func main() {
	s := "()()[(])"
	s = "()"
	s = ""
	fmt.Println(s, "is valid?  ", isValid(s))
}
