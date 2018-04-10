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

//var ngramsd = func(agrams []string, bgrams []string) int {
//	common := 0
//	counter := make(map[string]int)
//	for _, g := range agrams {
//		counter[g]++
//	}
//	for _, g := range bgrams {
//		counter[g]++
//	}
//	for _, times := range counter {
//		if times > 1 {
//			common++
//		}
//	}
//	return len(agrams) + len(bgrams) - 2*common
//}

//var ngramsds2 = func(a string, b string, n int) int {
//	agrams := wngrams(a, n)
//	bgrams := wngrams(b, n)
//	return ngramsd(agrams, bgrams)
//}

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
	fmt.Println("START:", s)
	rs := NewRankedStrings(len(ngd.dictGrams))
	sgramsm := wngramsm(s, ngd.N)
	for word, gramsm := range ngd.dictGrams {
		rs.Put(word, ngramsd(sgramsm, gramsm))
		ss := ngramsd(sgramsm, gramsm)
		if ss < 7 {
			fmt.Println(word, ss)
		}
	}
	return rs
}
