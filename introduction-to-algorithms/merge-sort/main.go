package main

import "fmt"

func main() {
	unsortedList := []int{77, 4, 22, 1, 14, 18}
	fmt.Println(mergeSort(unsortedList))
}

func mergeSort(list []int) []int {
	switch len(list) {
	case 0:
		return list
	case 1:
		return list
	default:
		m := len(list) / 2
		left := list[:m]
		right := list[m:]
		return merge(mergeSort(left), mergeSort(right))
	}
}

func merge(left, right []int) []int {
	result := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
