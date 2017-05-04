package main

import "fmt"

// equal tests the deep equality of two slices

func main() {
	x1 := []string{"a", "b", "c"}
	y1 := []string{"a", "b", "c"}
	y2 := []string{"a", "b", "d"}

	fmt.Println(equal(x1, y1)) // true
	fmt.Println(equal(x1, y2)) // false
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
