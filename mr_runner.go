package main

import (
	"log"
	"time"
	"sync"
	"fmt"
	"runtime"
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

	m := &runtime.MemStats{}

	runtime.ReadMemStats(m)
	println(m.Alloc / 1000 / 1000)
	t := time.Now()
	method.Prepare(am)
	println("Prepare: " + time.Since(t).String())

	runtime.ReadMemStats(m)
	println(m.Alloc / 1000 / 1000)
	counter.Start()
	var wg sync.WaitGroup
	var stepWg sync.WaitGroup
	step := 512
	mark := 0
	for i, s := range am.misspells {
		wg.Add(1)
		stepWg.Add(1)
		go func(i int, s string) {
			start := time.Now()
			rc := method.Match(am.dict, s)
			rc.Sort()
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
	runtime.ReadMemStats(m)
	println(m.Alloc / 1000 / 1000)
	//go func() {
	//	for {
	//		runtime.ReadMemStats(m)
	//		println(m.Alloc/1000/1000)
	//runtime.GC()
	//time.Sleep(time.Second * 16)
	//}
	//}()
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
