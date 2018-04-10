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
	//for i, v := range r.List {
	//	if v.R <= 3 {
	//		println(strconv.Itoa(i) + ") " + v.S + " " + strconv.Itoa(v.R))
	//	}
	//}
	r.List = append(make([]RankedString, 0, i), r.List[:i]...)

}

func (r *RankedStrings) TopIndex(limit int) int {
	if !r.Sorted {
		r.Sort()
	}
	var max int
	times := 0
	for i, rs := range r.List {
		if i == 0 {
			max = rs.R
			times++
		} else if rs.R > max {
			max = rs.R
			times++
		}
		if times > limit {
			return i
		}
	}
	return len(r.List)
}

func (r *RankedStrings) Top(limit int) []string {
	i := r.TopIndex(limit)
	result := make([]string, 0, i)
	for li, rs := range r.List {
		if li < i {
			result = append(result, rs.S)
		}
	}
	return result
}
