package main

func main() {
	numberOfSteps(10)
}

func numberOfSteps(num int) int {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
	return num
}
