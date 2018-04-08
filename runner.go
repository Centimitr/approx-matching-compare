package main

import (
	"sync"
	"log"
	"time"
	"fmt"
)

type Dict struct {
	List    []string
	Mapping map[string]struct{}
}

func NewDictFromFile(filename string) Dict {
	lines := ReadFileAsLines(filename)
	d := Dict{
		lines,
		make(map[string]struct{}, len(lines)),
	}
	for _, word := range d.List {
		d.Mapping[word] = struct{}{}
	}
	return d
}

func (d Dict) Has(s string) bool {
	if _, ok := d.Mapping[s]; ok {
		return true
	}
	return false
}

type ApproxMatchMethod interface {
	Match(Dict, string) (candidates []string, time int)
}

type ApproxMatchRecord struct {
	Method     string
	Parameter  string
	StartTime  int
	Candidates [][]string
	Times      []int

	// stat
	Hits      []bool  `json:",omitempty"`
	Precision float64 `json:",omitempty"`
	Recall    float64 `json:",omitempty"`
	FMeasure  float64 `json:",omitempty"`

	TotalTime int `json:",omitempty"`
	AvgTime   int `json:",omitempty"`
	MinTime   int `json:",omitempty"`
	MaxTime   int `json:",omitempty"`
	TimeDrop  int `json:",omitempty"`

	TimeCmp     float64 `json:",omitempty"`
	MinTimeCmp  float64 `json:",omitempty"`
	MaxTimeCmp  float64 `json:",omitempty"`
	TimeDropCmp float64 `json:",omitempty"`
}

func (r ApproxMatchRecord) Stat() {
}

type ApproxMatchTask struct {
	MisspellsPath string              `json:"misspells"`
	CorrectsPath  string              `json:"corrects"`
	DictPath      string              `json:"dict"`
	Records       []ApproxMatchRecord `json:"records,omitempty"`
}

func (t ApproxMatchTask) Stat() {
	var wg sync.WaitGroup
	for _, r := range t.Records {
		wg.Add(1)
		go func() {
			r.Stat()
			wg.Done()
		}()
	}
	wg.Wait()
}

type ApproxMatchRunner struct {
	task      ApproxMatchTask
	dict      Dict
	misspells []string
	corrects  []string
}

func NewApproxMatchRunner() ApproxMatchRunner {
	return ApproxMatchRunner{}
}

func (am ApproxMatchRunner) Load(filename string) {
	e := ReadJSON(filename, &am.task)
	if e != nil {
		log.Fatal("Read task json failed.")
	}
	am.dict = NewDictFromFile(am.task.DictPath)
	am.misspells = ReadFileAsLines(am.task.MisspellsPath)
	am.corrects = ReadFileAsLines(am.task.CorrectsPath)
}

func (am ApproxMatchRunner) Save(filename string) {
	e := WriteJSON(filename, am.task)
	if e != nil {
		log.Fatal("Write task json failed.")
	}
}

func (am ApproxMatchRunner) Run(method ApproxMatchMethod, limits []int, note string) {
	startTime := time.Now().UnixNano()
	methodName := GetFunctionName(method)
	candidates := make([][]string, len(am.misspells))
	times := make([]int, len(am.misspells))
	var wg sync.WaitGroup
	for i, s := range am.misspells {
		wg.Add(1)
		go func(i int, s string) {
			cs, t := method.Match(am.dict, s)
			candidates[i] = cs
			times[i] = t
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	for limit := range limits {
		r := ApproxMatchRecord{
			Method:     methodName,
			Candidates: candidates[:limit],
			Parameter:  fmt.Sprintf("(%s)-%d", note, limit),
			StartTime:  int(startTime),
			Times:      times,
		}
		am.task.Records = append(am.task.Records, r)
	}
}

func (am ApproxMatchRunner) Stat() {
	am.task.Stat()
}
