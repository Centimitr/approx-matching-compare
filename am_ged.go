package main

import "strings"

type GlobalEditDistance struct {
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
var editd = func(s, t string) int {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
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

func (ged GlobalEditDistance) Match(d Dict, s string) string {
	return minDistance(d.List, func(t string) int {
		return editd(s, t)
	})
}
