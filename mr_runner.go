package main

import (
	"log"
	"time"
	"sync"
	"fmt"
)

type ApproxMatchMethod interface {
	Match(Dict, string) RankedStrings
	Param() string
	Limits() []int
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
	am.task.Misspells = ReadFileAsLines(am.task.Path.Misspells)
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

func (am *ApproxMatchRunner) Run(method ApproxMatchMethod) {
	startTime := time.Now()
	methodName := GetStructName(method)
	rankedCandidates := make([]RankedStrings, len(am.misspells))
	times := make([]int, len(am.misspells))

	println("Start: " + methodName + " " + method.Param())
	counter := NewCounter(len(am.misspells))

	counter.Start()
	var wg sync.WaitGroup
	for i, s := range am.misspells {
		wg.Add(1)
		go func(i int, s string) {
			start := time.Now()
			rankedCandidates[i] = method.Match(am.dict, s)
			times[i] = int(time.Since(start))
			counter.Add()
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	counter.Finish()
	for _, limit := range method.Limits() {
		r := ApproxMatchRecord{
			Method: methodName,
			Candidates: func() [][]string {
				result := make([][]string, len(rankedCandidates))
				for i, rc := range rankedCandidates {
					result[i] = rc.Top(limit)
				}
				return result
			}(),
			Parameter: fmt.Sprintf("(%s)-%d", method.Param(), limit),
			StartTime: int(startTime.UnixNano()),
			Times:     times,
		}
		am.task.Records = append(am.task.Records, r)
	}

	println("Complete: " + time.Since(startTime).String() + "\n")
}

func (am *ApproxMatchRunner) Stat() {
	am.task.Stat(am.corrects)
}
