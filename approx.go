package main

import "strings"

// Neighbourhood Search
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

// N-Gram Distance
type NGramDistance struct {
	N int
}

var ngrams = func(s string, n int) (grams []string) {
	for i := 0; i <= len(s)-n; i++ {
		grams = append(grams, s[i:i+n])
	}
	return
}

var ngramsd = func(a string, b string, n int) int {
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

func (ngd NGramDistance) Match(d Dict, s string) string {
	//minD := 0
	//minI := 0
	//for i, word := range dict.List {
	//	d := ngramsd(s, word, ngd.N)
	//	if i == 0 {
	//		minD = d
	//	} else if d < minD {
	//		minD = d
	//		minI = i
	//	}
	//}
	//return dict.List[minI]
	return minDistance(d.List, func(t string) int {
		return ngramsd(s, t, ngd.N)
	})
}

// Edit Distance
type EditDistance struct {
}

var min = func(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
var editd = func(s, t string) int {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	var leftCol []int
	for x := 0; x <= len(ts); x++ {
		col := make([]int, len(ss)+1)
		for y := 0; y <= len(ss); y++ {
			switch {
			case x == 0:
				col[y] = y
			case y == 0:
				col[y] = x
			case ss[y-1] != ts[x-1]:
				col[y]++
				fallthrough
			default:
				col[y] += min(leftCol[y], leftCol[y-1], col[y-1])
			}
		}
		leftCol = col
	}
	return leftCol[len(ss)]
}

func (ged EditDistance) Match(d Dict, s string) string {
	return minDistance(d.List, func(t string) int {
		return editd(s, t)
	})
}

// util
func minDistance(list []string, handle func(t string) int) string {
	minD := 0
	minI := 0
	for i, t := range list {
		d := handle(t)
		if i == 0 {
			minD = d
		} else if d < minD {
			minD = d
			minI = i
		}
	}
	return list[minI]
}
