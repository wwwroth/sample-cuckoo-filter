package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func generateKeys(numKeys int) {

	filePath := os.Getenv("TMP_KEY_FILE")

	// Attempt to open the file for writing. The flags indicate that the file should be:
	// - Opened with write-only access (os.O_WRONLY)
	// - Created if it does not exist (os.O_CREATE)
	// - Truncated to zero length if it already exists, effectively clearing its contents (os.O_TRUNC)
	// The file mode (0644) sets the file permissions to be readable and writable by the owner, and readable by others.
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Could not open or create temp file to hold keys.")
	}

	// Ensure the file is closed when the function execution is completed.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Unable to close temp file.")
		}
	}(file)

	for i := 0; i < numKeys; i++ {

		// Convert the current integer to a byte slice for hashing.
		strBytes := []byte(strconv.Itoa(i))

		// Generate a SHA-256 hash of the byte slice.
		hashStr := sha256.Sum256(strBytes)

		// Format the hash as a hexadecimal string and append a newline character for file writing.
		line := fmt.Sprintf("%x\n", hashStr)

		// Write the hash string to the file.
		_, err := file.WriteString(line)
		if err != nil {
			msg := fmt.Sprintf("Failed writing to file: %s", err)
			panic(msg)
		}
	}
}
