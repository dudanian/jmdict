package main

import (
	"flag"
	"fmt"

	"github.com/dudanian/kanjisho/jmdict"
)

func main() {
	toUpdate := flag.Bool("update", false, "update the dict file")
	flag.Parse()

	if *toUpdate {
		fmt.Println("Updating the dictionary file")
		jmdict.Download()
	}

	if !*toUpdate {
		fmt.Println("Not doing anything...")
		fmt.Println("Select a command:")
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
	}
}
