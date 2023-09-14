package codesignal

import "fmt"

func centuryFromYear(year int) int {

	fmt.Println(((year - 1) / 100) + 1)

	return 0
}

// given 1935
// 1935 - 1
// 1934
// 1934 / 100
// 19.34
// 19.34 + 1
// 20
