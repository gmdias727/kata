package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	lista := make([]string, n)

	for i := range lista {
		value := i + 1
		if value%3 == 0 && value%5 == 0 {
			lista[i] = "FizzBuzz"
		} else if value%3 == 0 {
			lista[i] = "Fizz"
		} else if value%5 == 0 {
			lista[i] = "Buzz"
		} else {
			lista[i] = strconv.Itoa(value)
		}
	}
	return lista
}
