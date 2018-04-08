package main

import (
	"time"
	"fmt"
)

func NewTimer(count int, interval time.Duration) Timer {
	return Timer{N: count, Interval: interval}
}

type Timer struct {
	N        int
	Interval time.Duration
	i        int
	start    time.Time
	finish   bool
}

func (t *Timer) Add() {
	t.i ++
}

func (t *Timer) Print() {
	since := time.Since(t.start)
	fmt.Printf("\r%02d:%02d - %d/%d", int(since.Minutes()), int(since.Seconds())-int(since.Minutes())*60, t.i, t.N)
}

func (t *Timer) Start() {
	t.start = time.Now()
	t.i = 0
	t.finish = false
	go func() {
		for {
			if t.i >= t.N || t.finish {
				break
			}
			t.Print()
			time.Sleep(t.Interval)
		}
	}()
}

func (t *Timer) Finish() {
	t.Print()
	fmt.Println()
	t.finish = true
}
