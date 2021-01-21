package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nihaals/bottom-go/bottom"
)

func main() {
	decode := flag.Bool("d", false, "Whether to decode")
	flag.Parse()
	input := flag.Arg(0)
	if input == "" {
		fmt.Println("No input given")
		os.Exit(1)
	}
	if *decode {
		fmt.Println(bottom.Decode(input))
	} else {
		fmt.Println(bottom.Encode(input))
	}
}
