package main

import (
	"time"
	"strconv"
	"sync"
)

func NewCounter(count int) Counter {
	return Counter{N: count, RefreshInterval: time.Millisecond * 250, RefreshOnAdd: true}
}

type Counter struct {
	N               int
	RefreshInterval time.Duration
	RefreshOnAdd    bool
	i               int
	start           time.Time
	finish          bool
	wg              sync.WaitGroup
}

func (c *Counter) Add() {
	c.i ++
	c.wg.Done()
}

func (c *Counter) Print() {
	print("\r" + time.Since(c.start).String() + " - " + strconv.Itoa(c.i) + "/" + strconv.Itoa(c.N))
	//fmt.Printf("\r%02d:%02d - %d/%d", int(since.Minutes()),, t.i, t.N)
}

func (c *Counter) Start() {
	c.wg.Add(c.N)
	c.start = time.Now()
	c.i = 0
	c.finish = false
	go func() {
		for {
			if c.i >= c.N || c.finish {
				break
			}
			c.Print()
			time.Sleep(c.RefreshInterval)
		}
	}()
}

func (c *Counter) Finish() {
	c.finish = true
	c.wg.Wait()
	c.Print()
	println()
}
