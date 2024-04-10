package main

import "fmt"

func main() {
	// 2D array/slice
	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	maximumWealth(accounts)
}

func maximumWealth(accounts [][]int) {
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		fmt.Printf("Sum of row %d = %d\n", i+1, sum)
	}
}
