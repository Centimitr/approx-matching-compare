package main

type DirectMatch struct {
}

func (dm *DirectMatch) Name() string {
	return "DirectMatch"
}

func (dm *DirectMatch) Prepare(runner *ApproxMatchRunner) {
}

func (dm *DirectMatch) Step() int {
	return 4096
}

func (dm *DirectMatch) Match(d Dict, s string) RankedStrings {
	rs := NewRankedStrings(0)
	if d.Has(s) {
		rs.Put(s, 0)
	}
	return rs
}
