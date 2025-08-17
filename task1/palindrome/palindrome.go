package main

import (
	"fmt"
	"strconv"
)

func isPalindrome1(x int) bool {
	nums := strconv.Itoa(x)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reverse := 0

	for x > 0 {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	return original == reverse
}

func main() {
	for i := 101; i < 130; i++ {
		isPalindrome1 := isPalindrome1(i)
		if isPalindrome1 {
			fmt.Printf("num:%d is palindrome by isPalindrome1 func\n", i)
		}

		isPalindrome2 := isPalindrome2(i)
		if isPalindrome2 {
			fmt.Printf("num:%d is palindrome by isPalindrome2 func\n", i)
		}
	}

}
