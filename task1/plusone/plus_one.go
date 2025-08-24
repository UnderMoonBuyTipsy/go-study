package main

import "fmt"

func plusOne(digits []int) []int {
	digitsTemp := make([]int, len(digits))
	copy(digitsTemp, digits)
	for i := len(digitsTemp) - 1; i >= 0; i-- {
		digitsTemp[i]++
		digitsTemp[i] %= 10
		if digitsTemp[i] != 0 {
			return digitsTemp
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	digitsArr := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
	}
	for _, digits := range digitsArr {
		fmt.Println(digits, " plus one = ", plusOne(digits))
	}
}
