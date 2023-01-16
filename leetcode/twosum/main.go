package main

import "fmt"

func twoSum(nums []int, target int) []int {

	len := len(nums)

	for i := 0; i < len; i++ {

		check := nums[i]
		for j := i + 1; j < len; j++ {
			if check+nums[j] == target {
				var result []int
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return nil
}

func main() {
	// answer := twoSum()
	// func

	aa := make([]int)

	fmt.Println(aa)
}
