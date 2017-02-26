package main

import (
	"fmt"
	"os"
	"strconv"

	"exercises-go/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "kmm: %v\n", err)
			os.Exit(1)
		}
		lbs := unitconv.Pounds(w)
		kg := unitconv.Kilograms(w)
		fmt.Printf("%s = %s, %s = %s\n", lbs, unitconv.LbsToKg(lbs), kg, unitconv.KgToLbs(kg))
	}
}
