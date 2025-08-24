package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for i, c := range s0 {
		for _, str := range strs {
			if i == len(str) || str[i] != byte(c) {
				return s0[:i]
			}
		}
	}
	return s0
}

func main() {
	strsArr := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"dog", "door", "dot"},
	}
	for _, strs := range strsArr {
		fmt.Println(strs, "LongestCommonPrefix:", longestCommonPrefix(strs))
	}
}
