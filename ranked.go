package main

import (
	"sort"
	"strconv"
)

type RankedString struct {
	S string
	R int
}

func NewRankedStrings(cap int) RankedStrings {
	return RankedStrings{
		List: make([]RankedString, cap),
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

func (r *RankedStrings) Top(limit int) []string {
	if !r.Sorted {
		r.Sort()
	}
	//for _, v := range r.List {
	//	fmt.Printf("STR: %s - %d\n", v.S, v.R)
	//}
	print("LIMIT: " + strconv.Itoa(limit))
	result := make([]string, 0)
	n := 0
	var previous int
	for _, rs := range r.List {
		if rs.R != previous {
			n++
		}
		if n > limit {
			return result
		}
		result = append(result, rs.S)
		previous = rs.R
	}
	//fmt.Println(len(result), result[0])
	return result
}
