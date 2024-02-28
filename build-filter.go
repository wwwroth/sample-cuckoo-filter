package main

import (
	"bufio"
	cuckoo "github.com/panmari/cuckoofilter"
	"os"
)

func buildFilter() *cuckoo.Filter {

	file, err := os.Open(os.Getenv("TMP_KEY_FILE"))
	if err != nil {
		panic("Error opening temp key file.")
	}

	// Ensure the file is closed when the function execution is completed.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var keys []string

	// Create a new Scanner to read the file line by line
	// Read each line of the file using scanner.Scan in a loop.
	// For each line read, append it to the keys slice.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys, scanner.Text())
	}

	// After finishing the loop, check if there were any errors during scanning.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Initialize a new filter in memory twice the size of the existing keys
	cuckooFilter := cuckoo.NewFilter(uint(len(keys)) * 2)

	// Insert keys into filter
	for _, value := range keys {
		cuckooFilter.Insert([]byte(value))
	}

	return cuckooFilter
}
