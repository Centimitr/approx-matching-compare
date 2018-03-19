package main

import "strings"

type NeighbourhoodSearch struct {
	K int
}

var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
var search = func(d Dict, s string, alphabet []string) (word string, candidates []string) {
	//alen := len(alphabet)
	slen := len(s)
	//candidatesN := slen + slen*alen + (slen+1)*alen
	//candidates = make([]string, candidatesN)
	for i := 0; i <= slen; i++ {
		//replacement
		if i != slen {
			for _, c := range alphabet {
				word = s[:i] + c + s[i+1:]
				candidates = append(candidates, word)
				if d.Has(word) {
					return
				}
			}
		}
		//deletion
		if i != slen {
			word = s[:i] + s[i+1:]
			candidates = append(candidates, word)
			if d.Has(word) {
				return
			}
		}
		//insertion
		for _, c := range alphabet {
			word = s[:i] + c + s[i:]
			candidates = append(candidates, word)
			if d.Has(word) {
				return
			}
		}
	}
	return
}

func (ns NeighbourhoodSearch) Match(d Dict, s string) string {
	candidates := []string{s}
	for i := 0; i < ns.K; i++ {
		var newCandidates []string
		for _, c := range candidates {
			if r, cs := search(d, c, alphabet); r != "" {
				return r
			} else {
				newCandidates = append(cs)
			}
		}
		candidates = newCandidates
	}
	return ""
}

type EditDistance struct {
}

func (ged EditDistance) Match(d Dict, s string) string {
	return ""
}

type NGramDistance struct {
	N        int
	cache    [][]string
	hasCache bool
}

var ngrams = func(s string, n int) (grams []string) {
	s = "#" + s + "#"
	for i := 0; i <= len(s)-n; i++ {
		grams = append(grams, s[i:i+n])
	}
	return
}

var common = func(agrams []string, bgrams []string) (common int) {
	counter := make(map[string]int)
	for _, g := range agrams {
		counter[g]++
	}
	for _, g := range bgrams {
		counter[g]++
	}
	for _, times := range counter {
		if times > 1 {
			common++
		}
	}
	return
}

func (ngd NGramDistance) Match(dict Dict, s string) string {
	if !ngd.hasCache {
		ngd.cache = make([][]string, len(dict.List))
		for i, word := range dict.List {
			ngd.cache[i] = ngrams(word, ngd.N)
		}
		ngd.hasCache = true
	}
	minD := 0
	minI := 0
	sgrams := ngrams(s, ngd.N)
	for i, word := range dict.List {
		d := len(word) - 2*common(sgrams, ngd.cache[i])
		if i == 0 {
			minD = d
		} else if d < minD {
			minD = d
			minI = i
		}
	}
	return dict.List[minI]
}
