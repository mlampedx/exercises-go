package main

import (
	"fmt"
	"os"
	"strconv"
)

// reverse reverses a slice of ints in place
// rotate rotates a slice a number of positions equal to an int provided to the cli

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5, 4, 3, 2, 1]

	for _, pos := range os.Args[1:] {
		s := []int{0, 1, 2, 3, 4, 5}
		i, err := strconv.Atoi(pos)
		if err != nil {
			fmt.Println("Value provided must be of type int")
			panic(err)
		}
		reverse(s[:i])
		reverse(s[i:])
		reverse(s)
		fmt.Println(s)
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
