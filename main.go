package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/dudanian/kanjisho/jmdict"
)

func browse() {
	r, err := jmdict.NewGzipReader()
	if err != nil {
		log.Fatal(err)
	}

	decoder := jmdict.NewDecoder(r)

	for {
		var entry jmdict.Entry
		err := decoder.Entry(&entry)

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", entry)
		fmt.Println("Enter to continue, Ctrl+c to exit.")
		fmt.Scanln()
	}
}

func main() {
	toUpdate := flag.Bool("update", false, "update the dict file")
	toBrowse := flag.Bool("browse", false, "step through the dict file")
	flag.Parse()

	if *toUpdate {
		fmt.Println("Updating the dictionary file")
		jmdict.Download()
	}

	if *toBrowse {
		browse()
	}

	if !*toUpdate || !*toBrowse {
		fmt.Println("Not doing anything...")
		fmt.Println("Select a command:")
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
	}
}
