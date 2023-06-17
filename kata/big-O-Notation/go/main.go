package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(N^2)
// O(log N)
// O(n log n)

// var theArray []int
// var arraySize int
// var itemsInArray int
var startTime int64
var endTime int64

func main() {
	testAlgo := NewBigONotation(1_000_000)
	testAlgo.generateRandomArray()

	testAlgo2 := NewBigONotation(10_000_000)
	testAlgo2.generateRandomArray()

	testAlgo3 := NewBigONotation(100_000_000)
	testAlgo3.generateRandomArray()
	// testAlgo.linearSearchForValue(50)
	testAlgo.linearSearchForValue(20)
	testAlgo2.linearSearchForValue(20)
	testAlgo3.linearSearchForValue(20)
	// fmt.Println(newArray.addItemToArray(5))
}

type BigONotation struct {
	theArray     []int
	arraySize    int
	itemsInArray int
}

func NewBigONotation(size int) *BigONotation {
	return &BigONotation{
		arraySize:    size,
		theArray:     make([]int, size),
		itemsInArray: 0,
	}
}

func (b *BigONotation) generateRandomArray() []int {
	// fmt.Println(b.arraySize)
	for i := 0; i < b.arraySize; i++ {
		b.theArray[i] = rand.Intn(1000) + 10
	}
	return b.theArray
}

// O(1)
func (b *BigONotation) addItemToArray(newItem int) []int {
	// Literalmente adiciona um elemento ao Array (Pika)
	return append(b.theArray, newItem)
}

// O(N)
func (b *BigONotation) linearSearchForValue(value int) {
	valueInArray := false
	indexsWithValues := ""

	// fmt.Println(b.arraySize)
	// fmt.Println(value)

	startTime = time.Now().UnixMilli()

	for i := 0; i < b.arraySize; i++ {
		if b.theArray[i] == value {
			valueInArray = true
			indexsWithValues += fmt.Sprint(i) + " "
		}
	}

	fmt.Println("Value found: ", valueInArray)
	// fmt.Println("Value found: ", indexsWithValues)

	endTime = time.Now().UnixMilli()
	fmt.Println("Linear Search Took ", fmt.Sprint(endTime-startTime))
}
