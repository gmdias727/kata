package main

import "fmt"

func main() {
	nums := []int32{1, 3, 5, 7, 9}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int32) []int32 {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
