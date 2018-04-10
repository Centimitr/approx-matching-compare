package main

import (
	"sort"
)

var LIMIT_1 = NewApproxMatchMethodLimits(1)
var LIMIT_2 = NewApproxMatchMethodLimits(2)
var LIMIT_1_2 = NewApproxMatchMethodLimits(1, 2)
var LIMIT_1_5 = NewApproxMatchMethodLimits(1, 5)
var LIMIT_1_5_10 = NewApproxMatchMethodLimits(1, 5, 10)

func NewApproxMatchMethodLimits(limits ...int) ApproxMatchMethodLimits {
	sort.Ints(limits)
	if len(limits) > 0 {
		return ApproxMatchMethodLimits{limits: limits}
	}
	return ApproxMatchMethodLimits{limits: []int{1}}
}

type ApproxMatchMethodLimits struct {
	limits []int
}

func (l *ApproxMatchMethodLimits) Max() int {
	return l.limits[len(l.limits)-1]
}

func (l *ApproxMatchMethodLimits) Limits() []int {
	if l.limits == nil {
		return []int{1}
	}
	return l.limits
}
