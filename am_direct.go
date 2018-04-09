package main

type DirectMatch struct {
}

func (dm DirectMatch) Param() string {
	return ""
}

func (dm DirectMatch) Limits() []int {
	return []int{1}
}

func (dm DirectMatch) Match(d Dict, s string) RankedStrings {
	rs := NewRankedStrings(1)
	if d.Has(s) {
		rs.Put(s, 0)
	}
	return rs
}
