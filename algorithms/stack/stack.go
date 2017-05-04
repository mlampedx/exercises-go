package main

import "fmt"

// stack implementation in go

func main() {
	s := []int{0}
	s = push(s, 1)
	s = push(s, 2)
	s = push(s, 3)
	s = push(s, 4)
	s = push(s, 5)
	s = pop(s)
	s = remove(s, 4)
	fmt.Printf("%v\n", s) // [0, 1, 2, 3]
}

func push(stack []int, i int) []int {
	return append(stack, i)
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}

func remove(stack []int, i int) []int {
	copy(stack[:i], stack[i+1:])
	return stack[:len(stack)-1]
}
