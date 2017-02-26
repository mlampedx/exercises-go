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
			fmt.Fprintf(os.Stderr, "kmm: %v\n", err)
			os.Exit(1)
		}
		km := unitconv.Kilometers(d)
		mi := unitconv.Miles(d)
		fmt.Printf("%s = %s, %s = %s\n", km, unitconv.KmToMi(km), mi, unitconv.MiToKm(mi))
	}
}
