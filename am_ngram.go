package main

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