package main

import "fmt"

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}
	for i, num := range nums {
		if j, ok := idx[target-num]; ok {
			return []int{i, j}
		}
		idx[num] = i
	}
	return nil
}

func main() {
	numsArr := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}
	targetArr := []int{9, 6, 6}
	for i, nums := range numsArr {
		fmt.Println("Nums:", nums, " sum is target:", targetArr[i], " index:", twoSum(nums, targetArr[i]))
	}
}
