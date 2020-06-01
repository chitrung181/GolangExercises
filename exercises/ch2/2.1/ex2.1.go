package main

import (
	"fmt"
	"os"
	"strconv"

	temp "exercises/ch2/2.1/temperature"
)

func main() {
	for _, arg := range os.Args[1:] {
		temperature, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		cel := temp.Celsius(temperature)
		fah := temp.CToF(cel)
		kel := temp.CToK(cel)

		fmt.Printf("%s = %s = %s\n", cel.ToString(), fah.ToString(), kel.ToString())
	}
}
