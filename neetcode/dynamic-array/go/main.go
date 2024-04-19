package main

import (
	"fmt"
)

type DynArr struct {
	arr      []int
	capacity int
}

func main() {
	dynarr := DynArr{}
	dynarr.DynamicArray(5)
	dynarr.set(0, 10)
	dynarr.set(1, 15)
	dynarr.set(2, 20)
	fmt.Printf("dynarr: %v\n", dynarr.arr)
}

func (d *DynArr) DynamicArray(capacity int) []int {
	if capacity <= 0 {
		panic("capacity needs to be greater than 0")
	}
	d.arr = make([]int, capacity)
	return d.arr
}

func (d *DynArr) get(i int) int {
	return d.arr[i]
}

func (d *DynArr) set(i int, n int) {
	if i >= 0 && i <= len(d.arr) {
		d.arr[i] = n
	} else {
		fmt.Printf("i = %v n = %v", i, n)
		fmt.Println("Deu ruim")
	}
}

func (d *DynArr) pushBack(n int) {
	d.arr = append(d.arr, n)
}

func (d *DynArr) resize() {
	newArr := make([]int, len(d.arr)*2)
	copy(newArr, d.arr)
	fmt.Println("New slice: ", newArr)
}

func (d *DynArr) getCapacity() int {
	return len(d.arr)
}
