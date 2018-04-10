package main

import "fmt"

// N-Gram Distance
type NGramDistance struct {
	N         int
	dictGrams map[string]map[string]int
}

func (ngd *NGramDistance) Name() string {
	return fmt.Sprintf("nGram(N=%d)", ngd.N)
}

var ngrams = func(s string, n int) (grams []string) {
	for i := 0; i <= len(s)-n; i++ {
		grams = append(grams, s[i:i+n])
	}
	return
}

// weighted head and tail
var wngrams = func(s string, n int) (grams []string) {
	return ngrams("#"+s+"#", n)
}

var wngramsm = func(s string, n int) map[string]int {
	grams := wngrams(s, n)
	m := make(map[string]int, len(grams))
	for _, g := range grams {
		m[g]++
	}
	return m
}

var ngramsd = func(agrams, bgrams map[string]int) int {
	if len(bgrams) < len(agrams) {
		bgrams, agrams = agrams, bgrams
	}
	common := 0
	for gram, count := range agrams {
		bcount := bgrams[gram]
		if count < bcount {
			common += count
		} else {
			common += bcount
		}
	}
	return len(agrams) + len(bgrams) - 2*common
}

func (ngd *NGramDistance) Prepare(runner *ApproxMatchRunner) {
	ngd.dictGrams = make(map[string]map[string]int, len(runner.dict.List))
	for _, word := range runner.dict.List {
		ngd.dictGrams[word] = wngramsm(word, ngd.N)
	}
}

func (ngd *NGramDistance) Match(d Dict, s string) RankedStrings {
	rs := NewRankedStrings(len(ngd.dictGrams))
	sgramsm := wngramsm(s, ngd.N)
	for word, gramsm := range ngd.dictGrams {
		rs.Put(word, ngramsd(sgramsm, gramsm))
	}
	return rs
}
