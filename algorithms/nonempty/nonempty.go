package main

import "fmt"

// nonempty and nonempty2 return a slice void of empty strings

func main() {
	s1 := []string{"a", "", "b", "", "c", ""}
	s2 := []string{"a", "", "b", "", "c", ""}
	s3 := nonempty(s1)
	s4 := nonempty2(s2)
	fmt.Printf("%q\n", s3) // ["a", "b", "c"]
	fmt.Printf("%q\n", s4) // ["a", "b", "c"]
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	// zero-length slice of original arg
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
