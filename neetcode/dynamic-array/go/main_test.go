package main

import "testing"

func TestDynamicArrayWithPositiveCapacity(t *testing.T) {
	d := DynArr{}
	capacity := 5
	result := d.DynamicArray(capacity)

	if len(result) != capacity {
		t.Errorf("Expected length of result to be %d, but got %d", capacity, len(result))
	}

	for _, v := range result {
		if v != 0 {
			t.Errorf("Expected all elements in result to be 0, but got %d", v)
		}
	}
}
