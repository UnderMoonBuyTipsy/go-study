package main

import "fmt"

func removeDuplicates(nums []int) int {
	slowIndex := 1
	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != nums[fastIndex-1] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func main() {
	numsArr := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	for _, nums := range numsArr {
		fmt.Println(nums, " Remove Duplicates length=", removeDuplicates(nums))
	}
}
