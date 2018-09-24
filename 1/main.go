package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if i != j {
				if a+b == target {
					ret := make([]int, 2, 2)
					ret[0] = i
					ret[1] = j
					return ret
				}
			}
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
