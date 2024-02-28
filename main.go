package main

import (
	"flag"
)

func main() {

	actionPtr := flag.String("action", "", "What action would you like to perform?")
	numKeysPtr := flag.Int("numKeys", 1024, "How many keys would you like to generate?")
	flag.Parse()

	// Dereference the pointers to get the actual values of the command-line arguments.
	action := *actionPtr
	numKeys := *numKeysPtr

	switch action {
	case "generateKeys":
		generateKeys(numKeys)
	case "buildFilter":
		cuckooFilter := buildFilter()
		startServer(cuckooFilter)
	}
}
