package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFileAsLines(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(string(content))
}
