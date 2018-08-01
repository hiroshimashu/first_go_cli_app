package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lineNumber := flag.Int("n", 10, "How many lines to get.")
	flag.Parse()

	result := fmt.Sprintf("line numbers: %d", *lineNumber)
	fmt.Println(result)
}

// Parse text file.
func parseLines(filePath string) ([]string, error) {
	input, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var results []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		log.Print(strconv.Quote(scanner.Text()))
	}
	if scanner.Err() != nil {
		err := scanner.Err()
		return nil, err
	}
}
