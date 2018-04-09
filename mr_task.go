package main

import "sync"

type ApproxMatchTask struct {
	MisspellsPath string              `json:"misspells"`
	CorrectsPath  string              `json:"corrects"`
	DictPath      string              `json:"dictionary"`
	Records       []ApproxMatchRecord `json:"records,omitempty"`
}

func (t ApproxMatchTask) Stat(corrects []string) {
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
