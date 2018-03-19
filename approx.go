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
	N int
}

var ngrams = func(s string, n int) (grams []string) {
	for i := 0; i <= len(s)-n; i++ {
		grams = append(grams, s[i:i+n])
	}
	return
}

var ngramd = func(a string, b string, n int) int {
	agrams := ngrams("#"+a+"#", n)
	bgrams := ngrams("#"+b+"#", n)
	common := 0
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
	return len(agrams) + len(bgrams) - 2*common
}

func (ngd NGramDistance) Match(dict Dict, s string) string {
	minD := 0
	minI := 0
	//dictGrams := make([][]string, len(dict.List))
	// 4823
	for i, word := range dict.List {
		d := ngramd(s, word, ngd.N)
		if i == 0 {
			minD = d
		} else if d < minD {
			minI = i
		}
	}
	return dict.List[minI]
}
