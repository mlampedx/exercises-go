package main

import "fmt"

// equal-map checks whether two maps are deeply equal or not

func main() {
	var m1 = map[string]int{"A": 0, "B": 1, "C": 3}
	var m2 = map[string]int{"A": 0, "B": 1, "C": 3}
	b := equal(m1, m2)
	fmt.Println(b) // true
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
