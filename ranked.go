package main

import (
	"sort"
)

type RankedString struct {
	S string
	R int
}

func NewRankedStrings(cap int) RankedStrings {
	return RankedStrings{
		List: make([]RankedString, 0, cap),
	}
}

type RankedStrings struct {
	List   []RankedString
	Sorted bool
}

func (r *RankedStrings) Has(s string) bool {
	for _, rs := range r.List {
		if rs.S == s {
			return true
		}
	}
	return false
}

func (r *RankedStrings) Put(s string, rank int) {
	r.List = append(r.List, RankedString{s, rank})
}

func (r *RankedStrings) Sort() {
	sort.Slice(r.List, func(i, j int) bool {
		return r.List[i].R < r.List[j].R
	})
	r.Sorted = true
}
func (r *RankedStrings) Shrink(limit int) {
	i := r.TopIndex(limit)
	r.List = r.List[:i]
}

func (r *RankedStrings) TopIndex(limit int) int {
	if !r.Sorted {
		r.Sort()
	}
	n := 0
	var previous int
	for i, rs := range r.List {
		if i == 0 || rs.R != previous {
			n++
		}
		if n > limit {
			return n
		}
		previous = rs.R
	}
	return len(r.List)
}

func (r *RankedStrings) Top(limit int) []string {
	i := r.TopIndex(limit)
	//nn := 0
	//for _, v := range r.List {
	//	if v.R < 3 {
	//		nn++
	//		fmt.Printf("STR: %s - %d\n", v.S, v.R)
	//	}
	//}
	//fmt.Println(nn)
	result := make([]string, 0, i)
	for _, rs := range r.List {
		result = append(result, rs.S)
	}
	return result
}
