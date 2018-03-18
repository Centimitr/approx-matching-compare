package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

func ReadFileAsLines(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(string(content))
}

func PrintLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
