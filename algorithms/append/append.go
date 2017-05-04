package main

import "fmt"

// appendInt is a different implementation of the built-in append function

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		if i == 9 {
			y = appendInt(x, 9, 10, 11, 12)
		}

		fmt.Printf("%d	cap=%d\t%v\n", i, cap(y), y)

		// 0	cap=1	[0]
		// 1	cap=2	[0 1]
		// 2	cap=4	[0 1 2]
		// 3	cap=4	[0 1 2 3]
		// 4	cap=8	[0 1 2 3 4]
		// 5	cap=8	[0 1 2 3 4 5]
		// 6	cap=8	[0 1 2 3 4 5 6]
		// 7	cap=8	[0 1 2 3 4 5 6 7]
		// 8	cap=16	[0 1 2 3 4 5 6 7 8]
		// 9	cap=16	[0 1 2 3 4 5 6 7 8 9 10 11 12]

		x = y
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	var zlen int
	if len(y) > 1 {
		zlen = len(x) + len(y)
	} else {
		zlen = len(x) + 1
	}
	if zlen <= cap(x) {
		// there is room to grow; extend the slice
		z = x[:zlen]
	} else {
		// there is insufficient space; allocate new arr
		// grow by doubling for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		// copy contents of x into z
		copy(z, x)
	}
	if len(y) > 1 {
		copy(z[len(x):], y)
	} else {
		z[len(x)] = y[0]
	}
	return z
}
