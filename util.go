package main

import (
	"io/ioutil"
	"log"
	"strings"
	"encoding/json"
	"os"
	"reflect"
)

func ReadFileAsLines(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(string(content))
}

func ReadJSON(filename string, v interface{}) error {
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		return e
	}
	e = json.Unmarshal(b, v)
	return e
}

func WriteJSON(filename string, v interface{}) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	e = ioutil.WriteFile(filename, b, os.ModePerm)
	return e
}
func GetStructName(v interface{}) string {
	return reflect.TypeOf(v).String()
}
