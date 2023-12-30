package main

import (
	"fmt"
	"hash/crc32"
	"os"
)

type result struct {
	checksum uint32
	err      error
}

func calculateChecksum(filePath string, resultChan chan<- result) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error, failed to read file", filePath)
		resultChan <- result{checksum: 0, err: err}
	}

	// Create a new CRC-32 hasher
	hasher := crc32.NewIEEE()

	// Write the file content to the hasher
	_, err = hasher.Write(content)
	if err != nil {
		fmt.Println("Error, failed to write data in to hasher")
		resultChan <- result{checksum: 0, err: err}
	}

	// Get the CRC-32 checksum
	checksum := hasher.Sum32()

	resultChan <- result{checksum: checksum, err: nil}
}

func main() {
	filePath := "./file1.txt"

	// channel is inter-thread communication mechanism in go
	resultChan1 := make(chan result)

	// Start a go routine (new thread) that will calculate the checksum
	go calculateChecksum(filePath, resultChan1)

	// Get the calculation result
	result := <-resultChan1

	if result.err != nil {
		fmt.Printf("Error calculating checksum: %v\n", result.err)
		os.Exit(1)
	}

	fmt.Printf("Checksum for file %s: %08x\n", filePath, result.checksum)
}
