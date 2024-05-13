// https://github.com/soupi/3-C-Sorting-Algorithms/blob/master/InsertionSort.c
package main

import "fmt"

func main() {
	list := []int{7, 10, 1, 12, 50, 22, 14, 79}
	fmt.Println(insertionSort(list))
	// output: [1 7 10 12 14 22 50 79]
}

func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		j := i
		for j > 0 && list[j-1] > list[j] {
			swap(j-1, j, list)
			j -= 1
		}
	}
	return list
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
