package main

import (
	"sync"
)

type Dict struct {
	List    []string
	Mapping map[string]struct{}
}

func NewDictFromFile(filename string) Dict {
	lines := ReadFileAsLines(filename)
	d := Dict{
		lines,
		make(map[string]struct{}, len(lines)),
	}
	for _, word := range d.List {
		d.Mapping[word] = struct{}{}
	}
	return d
}

func (d Dict) Has(s string) bool {
	if _, ok := d.Mapping[s]; ok {
		return true
	}
	return false
}

type ApproxMatchMethod interface {
	Match(Dict, string) string
}

func (d Dict) MultiApproxMatch(words []string, method ApproxMatchMethod) []string {
	results := make([]string, len(words))
	var wg sync.WaitGroup
	for i, s := range words {
		wg.Add(1)
		go func(i int, s string) {
			results[i] = method.Match(d, s)
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	return results
}
