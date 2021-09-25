package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anders-14/spraok/lexer"
)

func readFromFile(filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return string(contents)
}

func main() {
	l := lexer.New(readFromFile("./test.spraok"))

	for !l.Done() {
		fmt.Println(l.NextToken().HumanReadable())
	}
}
