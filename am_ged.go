package main

type GlobalEditDistance struct {
	dictChars map[string][]rune
}

func (ged *GlobalEditDistance) Name() string {
	return "GED"
}

func (ged *GlobalEditDistance) Step() int {
	return 128
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
var editd = func(ss, ts []rune) int {
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

func (ged *GlobalEditDistance) Prepare(runner *ApproxMatchRunner) {
	ged.dictChars = make(map[string][]rune, len(runner.dict.List))
	for _, word := range runner.dict.List {
		ged.dictChars[word] = []rune(word)
	}
}

func (ged *GlobalEditDistance) Match(d Dict, s string) RankedStrings {
	rs := NewRankedStrings(len(ged.dictChars))
	srunes := []rune(s)
	for word, runes := range ged.dictChars {
		rs.Put(word, editd(runes, srunes))
	}
	return rs
}
