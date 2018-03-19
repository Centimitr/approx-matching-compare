package main

import (
	"time"
)

func main() {
	dict := NewDictFromFile("dictionary.txt")
	misspells := ReadFileAsLines("misspell.txt")[:10]
	corrects := ReadFileAsLines("correct.txt")[:10]
	t := time.Now()
	// 92 624, 12%, 243
	//results := dict.MultiApproxMatch(misspells, NeighbourhoodSearch{K: 1})
	// 99 617, 13%, 4090569
	results := dict.MultiApproxMatch(misspells, NGramDistance{N: 2})
	//
	//results := dict.MultiApproxMatch(misspells, EditDistance{})
	success := 0
	fail := 0
	for i, r := range results {
		m := misspells[i]
		c := corrects[i]
		println(m, r, c)
		if r == c {
			success++
		} else {
			fail ++
		}
	}
	println(success, fail)
	println(success*100/(success+fail), "%")
	println("TIME:", time.Since(t).Nanoseconds()/1e5, "ms")
}
