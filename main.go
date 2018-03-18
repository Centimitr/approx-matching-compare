package main

import (
	"time"
)

func main() {
	dict := NewDictFromFile("dictionary.txt")
	misspells := ReadFileAsLines("misspell.txt")
	corrects := ReadFileAsLines("correct.txt")
	t := time.Now()
	ns := NeighbourhoodSearch{K: 1}
	results := dict.MultiApproxMatch(misspells, ns)
	success := 0
	fail := 0
	for i, r := range results {
		//m := misspells[i]
		c := corrects[i]
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
