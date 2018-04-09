package main

import (
	"io/ioutil"
	"log"
	"strings"
	"encoding/json"
	"os"
	"reflect"
	"sort"
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

type RankedString struct {
	S string
	R int
}

func NewRankedStrings(cap int) RankedStrings {
	return RankedStrings{
		List: make([]RankedString, cap),
	}
}

type RankedStrings struct {
	List   []RankedString
	Sorted bool
}

func (r *RankedStrings) Put(s string, rank int) {
	r.List = append(r.List, RankedString{s, rank})
}

func (r *RankedStrings) Sort() {
	sort.Slice(r.List, func(i, j int) bool {
		return r.List[i].R < r.List[j].R
	})
	r.Sorted = true
}

func (r *RankedStrings) Top(limit int) (result []string) {
	if !r.Sorted {
		r.Sort()
	}
	n := 0
	var previous int
	for _, rs := range r.List {
		if rs.R != previous {
			n++
		}
		if n > limit {
			return
		}
		result = append(result, rs.S)
	}
	return
}
