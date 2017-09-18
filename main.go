package main

import (
	"fmt"
	"os"
	"strconv"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: thou <number> [... <number>]")
		os.Exit(1)
	}

	for _, number := range os.Args[1:] {
		actualNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Error converting %s to number: %v\n", number, err)
			os.Exit(1)
		}

		numberWithCommas := humanize.Comma(int64(actualNumber))
		fmt.Printf("%s -> %s\n", number, numberWithCommas)
	}
}
