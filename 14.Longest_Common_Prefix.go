package main

import (
	"bytes"
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var buf bytes.Buffer
	for i := 0; ; i++ {
		set := make(map[string]bool)
		if len(strs) == 0 {
			return buf.String()
		}
		for j := 0; j < len(strs); j++ {
			s := strs[j]
			if len(s) <= i {
				return buf.String()
			}
			set[string(s[i])] = true
			fmt.Println("set", set)
			if len(set) > 1 {
				return buf.String()
			}
		}
		buf.WriteString(string(strs[0][i]))
	}

	return buf.String()
}

func main() {
	strs := []string{"ab", "aba", "abc"}
	//strs := []string{}
	fmt.Println("outcome:", longestCommonPrefix(strs))
}
