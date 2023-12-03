package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 1 Problem")
	input := ReadFile("input.txt")
	for _, line := range input {
		fmt.Println(line)
	}
}

/// UTILS /// 
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


