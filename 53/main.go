package main

import "fmt"

func maxSubArray(nums []int) int {
	maxsum := -9223372036854775808
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > maxsum {
			maxsum = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}

	return maxsum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
