package main

import (
	"fmt"
	"greetings"
	"math"
	"math/rand"
	"reflect"
)

/*
	Equivalente
	func add(x int, y int)
	func add(x, y int)
*/

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "World")

	fmt.Println(greetings.Hello("Gabriel"))
	fmt.Println("My favorite number is: ", rand.Intn(100))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(reflect.TypeOf(math.Acos(0.987654321)))
	fmt.Println(add(27, 3))
	fmt.Println(a, b)
	// fmt.Println(morestrings.ReverseRunes("Teste"))
	// message := greetings.Hello("Gladys")
	// fmt.Println(message)
}
