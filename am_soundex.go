package main

import "fmt"

type Soundex struct {
	Cut           int
	dictSoundeces map[string]string
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

func (sd *Soundex) Name() string {
	return fmt.Sprintf("Soundex(Cut=%d)", sd.Cut)
}

func (sd *Soundex) Prepare(runner *ApproxMatchRunner) {
	sd.dictSoundeces = make(map[string]string, len(runner.dict.List))
	for _, word := range runner.dict.List {
		sd.dictSoundeces[word] = soundex(word, sd.Cut)
	}
}

func (sd *Soundex) Step() int {
	return 2048
}

func (sd *Soundex) Match(d Dict, s string) RankedStrings {
	rs := NewRankedStrings(0)
	ssd := soundex(s, sd.Cut)
	for word, wsd := range sd.dictSoundeces {
		if wsd == ssd {
			rs.Put(word, 0)
		}
	}
	return rs
}
