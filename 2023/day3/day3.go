package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 3 Problem")
	input := ReadFile("sample_input.txt")
	fmt.Println("Input", input)
	for _, line := range input {
		fmt.Println("Line", line)
	}
}

// / UTILS ///
// ReadFile reads a file and returns a slice of strings
func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		return nil
	}
	defer file.Close()

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}
