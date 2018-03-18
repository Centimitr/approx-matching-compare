package main

import "strings"

type NeighbourhoodSearch struct {
	K int
}

func (ns NeighbourhoodSearch) Match(d Dict, s string) string {
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	search := func(s string, alphabet []string) (word string, candidates []string) {
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
	kSearch := func(s string, k int) string {
		candidates := []string{s}
		for i := 0; i < k; i++ {
			var newCandidates []string
			for _, c := range candidates {
				if r, cs := search(c, alphabet); r != "" {
					return r
				} else {
					newCandidates = append(cs)
				}
			}
			candidates = newCandidates
		}
		return ""
	}
	return kSearch(s, ns.K)
}
