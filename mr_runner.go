package main

import (
	"log"
	"time"
	"sync"
	"fmt"
)

type ApproxMatchMethod interface {
	Prepare(*ApproxMatchRunner)
	Match(Dict, string) RankedStrings
	Name() string
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
	t = time.Now()
	for _, limit := range limits.Limits() {
		r := ApproxMatchRecord{
			Method: methodName,
			Candidates: func() [][]string {
				result := make([][]string, len(rankedCandidates))
				for i, rc := range rankedCandidates {
					go func(i int, rc RankedStrings, limit int) {
						result[i] = rc.Top(limit)
					}(i, rc, limit)
				}
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
}

func (am *ApproxMatchRunner) Stat() {
	am.task.Stat(am.corrects)
}
