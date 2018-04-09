package main

type Soundex struct {
	Cut int
}

func runesContains(rs []rune, tg rune) bool {
	for _, r := range rs {
		if r == tg {
			return true
		}
	}
	return false
}

func soundex(s string, cut int) string {
	var last rune
	var result []rune
	// s is an english string
	m := map[rune]int{
		'b': 1,
		'f': 1,
		'p': 1,
		'v': 1,
		'c': 2,
		'g': 2,
		'j': 2,
		'k': 2,
		'q': 2,
		's': 2,
		'x': 2,
		'z': 2,
		'd': 3,
		't': 3,
		'l': 4,
		'm': 5,
		'n': 5,
		'r': 6,
	}
	for i, r := range s {
		switch {
		case i == 0:
			result = append(result, r)
		case i == cut:
			break
		case last != r:
			result = append(result, rune(m[r]))
			last = r
		}
	}
	return string(result)
}

func (sd Soundex) Match(d Dict, s string) string {
	ssd := soundex(s, sd.Cut)
	for _, w := range d.List {
		if soundex(w, sd.Cut) == ssd {
			return w
		}
	}
	return ""
}
