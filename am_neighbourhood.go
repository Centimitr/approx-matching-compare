package main

import "strings"

type NeighbourhoodSearch struct {
	K int
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
//
//func (ns NeighbourhoodSearch) Match(d Dict, s string) RankedStrings {
//	elms := []string{s}
//	for i := 0; i < ns.K; i++ {
//		var newCandidates []string
//		for _, e := range elms {
//			cs, next := search(d, e, alphabet)
//
//			//if r, cs := search(d, c, alphabet); r != "" {
//			//	return r
//			//} else {
//			//	newCandidates = append(cs)
//			//}
//		}
//		candidates = newCandidates
//	}
//	return ""
//}
