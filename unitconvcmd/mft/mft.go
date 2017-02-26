package main

import (
	"fmt"
	"os"
	"strconv"

	"exercises-go/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		d, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mft: %v\n", err)
			os.Exit(1)
		}
		m := unitconv.Meters(d)
		ft := unitconv.Feet(d)
		fmt.Printf("%s = %s, %s = %s\n", m, unitconv.MToFt(m), ft, unitconv.FtToM(ft))
	}
}
