package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Open the file in append mode, or create it if it doesn't exist.
	// The 0644 is the file permission.
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Prepare the data to be written.
	data := fmt.Sprintf("Data added at: %s\n", time.Now().Format(time.RFC3339))

	// Write the data to the file.
	if _, err := f.WriteString(data); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data successfully written to test.txt")
}
