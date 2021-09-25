package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readLinesFromFile(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(contents), "\n")

	// Make sure to remove the empty line, that doesnt really exist,
	// behind the last newline character
	return lines[:len(lines)-1]
}

func main() {
	lines := readLinesFromFile("./test.spraak")
	fmt.Println(lines)
}
