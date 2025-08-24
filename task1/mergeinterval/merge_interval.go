package main

import (
	"fmt"
	"slices"
)

func mergeInterval(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p []int, q []int) int { return p[0] - q[0] })
	for _, interval := range intervals {
		m := len(ans)
		if m > 0 && interval[0] <= ans[m-1][1] {
			ans[m-1][1] = max(ans[m-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}
	return
}

func main() {
	intervalsArr := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
	}
	for _, intervals := range intervalsArr {
		ans := mergeInterval(intervals)
		fmt.Println("Intevals:", intervals, "merge result:", ans)
	}
}
