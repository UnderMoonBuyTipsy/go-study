package main

import "fmt"

func isValidBracket(s string) bool {
	stack := make([]rune, 0, 100)
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != bracketMap[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

func main() {
	strs := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, str := range strs {
		fmt.Printf("%q, valid bracket:%v\n", str, isValidBracket(str))
	}
}
