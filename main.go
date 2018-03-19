package main

import (
	"time"
)

func main() {
	dict := NewDictFromFile("dictionary.txt")
	misspells := ReadFileAsLines("misspell.txt")
	corrects := ReadFileAsLines("correct.txt")
	t := time.Now()
	// 92 624, 12%, 148
	results := dict.MultiApproxMatch(misspells, NeighbourhoodSearch{K: 1})
	// 99 617, 13%, 4,090,569
	//results := dict.MultiApproxMatch(misspells, NGramDistance{N: 2})
	// 116 600, 16%, 1,662,917
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
