package main

import "sync"

type ApproxMatchTaskPath struct {
	Dict      string `json:"dictionary"`
	Misspells string `json:"misspells"`
	Corrects  string `json:"corrects"`
}

type ApproxMatchTask struct {
	Path       ApproxMatchTaskPath `json:"path"`
	Misspells  []string            `json:"misspells"`
	Corrects   []string            `json:"corrects"`
	Records    []ApproxMatchRecord `json:"records,omitempty"`
	ProcessNum int                 `json:"processNum,omitempty"`
}

func (t *ApproxMatchTask) Stat(corrects []string) {
	var wg sync.WaitGroup
	for _, r := range t.Records {
		wg.Add(1)
		go func() {
			r.Stat(corrects)
			wg.Done()
		}()
	}
	wg.Wait()
}
