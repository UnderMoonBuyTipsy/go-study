package controlflow

func SingleNumber1(nums []int) int {
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		if v, exist := numMap[num]; exist {
			numMap[num] = v + 1
		} else {
			numMap[num] = 1
		}
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func SingleNumber2(nums []int) int {
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		numMap[num]++
	}
	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func SingleNumber3(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// func main() {
// 	nums1 := []int{1, 2, 1, 4, 2, 3, 4}

// 	num1 := SingleNumber1(nums1)

// 	num2 := SingleNumber2(nums1)

// 	fmt.Println("方法singleNumber1数组中只出现一次的数字为:", num1)

// 	fmt.Println("方法singleNumber2数组中只出现一次的数字为:", num2)
// }
