package main

import (
	"log"
	"time"
	"sync"
	"fmt"
	"runtime"
)

type ApproxMatchMethod interface {
	Name() string
	Prepare(*ApproxMatchRunner)
	Step() int
	Match(Dict, string) RankedStrings
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

func (am *ApproxMatchRunner) Load(filename string) {
	e := ReadJSON(filename, &am.task)
	if e != nil {
		log.Fatal("Read task json failed.")
	}
	am.dict = NewDictFromFile(am.task.Path.Dict)
	am.task.Dict = am.dict.List
	am.task.Misspells = ReadFileAsLines(am.task.Path.Misspells)
	if am.task.ProcessNum != 0 {
		am.task.Misspells = am.task.Misspells[:am.task.ProcessNum]
	}
	am.task.Corrects = ReadFileAsLines(am.task.Path.Corrects)
	am.misspells = am.task.Misspells
	am.corrects = am.task.Corrects
}

func (am *ApproxMatchRunner) Save(filename string) {
	e := WriteJSON(filename, am.task)
	if e != nil {
		log.Fatal("Write task json failed.")
	}
}

func (am *ApproxMatchRunner) Run(method ApproxMatchMethod, limits ApproxMatchMethodLimits) {
	startTime := time.Now()
	methodName := GetStructName(method)
	rankedCandidates := make([]RankedStrings, len(am.misspells))
	times := make([]int, len(am.misspells))

	println("Launch: " + method.Name())
	counter := NewCounter(len(am.misspells))

	t := time.Now()
	method.Prepare(am)
	println("Prepare: " + time.Since(t).String())

	counter.Start()
	var wg sync.WaitGroup
	var stepWg sync.WaitGroup
	step := method.Step()
	mark := 0
	for i, s := range am.misspells {
		wg.Add(1)
		stepWg.Add(1)
		go func(i int, s string) {
			start := time.Now()
			rc := method.Match(am.dict, s)
			rc.Shrink(limits.Max())
			rankedCandidates[i] = rc
			times[i] = int(time.Since(start))
			counter.Add()
			wg.Done()
			stepWg.Done()
		}(i, s)
		if i-mark > step {
			stepWg.Wait()
			mark = i
		}
	}
	wg.Wait()
	counter.Finish()
	t = time.Now()
	for _, limit := range limits.Limits() {
		r := ApproxMatchRecord{
			Method: methodName,
			Candidates: func() [][]string {
				result := make([][]string, len(rankedCandidates))
				var wg sync.WaitGroup
				for i, rc := range rankedCandidates {
					wg.Add(1)
					go func(i int, rc RankedStrings, limit int) {
						result[i] = rc.Top(limit)
						wg.Done()
					}(i, rc, limit)
				}
				wg.Wait()
				return result
			}(),
			Name:      fmt.Sprintf("%s-%d", method.Name(), limit),
			StartTime: int(startTime.UnixNano()),
			Times:     times,
		}
		am.task.Records = append(am.task.Records, r)
	}
	println("Complete: " + time.Since(t).String())
	println("Total: " + time.Since(startTime).String() + "\n")
	runtime.GC()
}

func (am *ApproxMatchRunner) Stat() {
	am.task.Stat(am.corrects)
}
