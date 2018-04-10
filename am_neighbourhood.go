package main

import (
	"strings"
	"fmt"
)

type NeighbourhoodSearch struct {
	K int
}

func (ns *NeighbourhoodSearch) Prepare(runner *ApproxMatchRunner) {
}

func (ns *NeighbourhoodSearch) Name() string {
	return fmt.Sprintf("Neighbour(K=%d)", ns.K)
}

var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
var search = func(d Dict, s string, alphabet []string) (candidates []string, next []string) {
	var word string
	//alen := len(alphabet)
	slen := len(s)
	//candidatesN := slen + slen*alen + (slen+1)*alen
	//candidates = make([]string, candidatesN)
	for i := 0; i <= slen; i++ {
		//replacement
		if i != slen {
			for _, c := range alphabet {
				word = s[:i] + c + s[i+1:]
				next = append(next, word)
				if d.Has(word) {
					candidates = append(candidates, word)
				}
			}
		}
		//deletion
		if i != slen {
			word = s[:i] + s[i+1:]
			next = append(next, word)
			if d.Has(word) {
				candidates = append(candidates, word)
			}
		}
		//insertion
		for _, c := range alphabet {
			word = s[:i] + c + s[i:]
			next = append(next, word)
			if d.Has(word) {
				candidates = append(candidates, word)
			}
		}
	}
	return
}

func (ns *NeighbourhoodSearch) Match(d Dict, s string) (rs RankedStrings) {
	seeds := []string{s}
	for i := 0; i < ns.K; i++ {
		var nextSeeds []string
		for _, seed := range seeds {
			cs, next := search(d, seed, alphabet)
			for _, c := range cs {
				if !rs.Has(c) {
					rs.Put(c, i+1)
				}
			}
			for _, s := range next {
				nextSeeds = append(nextSeeds, s)
			}
		}
		seeds = nextSeeds
	}
	rs.Sort()
	//for _, r := range rs.List {
	//	println(r.S, r.R)
	//}
	return
}
