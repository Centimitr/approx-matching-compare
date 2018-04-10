package main

import (
	"time"
	"os/exec"
	"net/http"
	"sync"
)

func main() {
	println("COMP90049 - Project 1")
	println("Student ID: 879849")
	println("Name: Xiao Shi")
	println()
	t := time.Now()

	am := NewApproxMatchRunner()
	am.Load("task.json").
		Run(&DirectMatch{}, LIMIT_1).
		Run(&NeighbourhoodSearch{K: 1}, LIMIT_1).
		Run(&NeighbourhoodSearch{K: 2}, LIMIT_2).
		Run(&NGramDistance{N: 2}, LIMIT_1).
		Run(&NGramDistance{N: 3}, LIMIT_1).
		Run(&NGramDistance{N: 4}, LIMIT_1).
		Run(&GlobalEditDistance{}, LIMIT_1_2).
		Run(&Soundex{Cut: 4}, LIMIT_1).
		Run(&Soundex{Cut: 8}, LIMIT_1).
		Stat().
		Save("result.json")

	since := time.Since(t)
	println("TIME: " + since.String())

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		http.ListenAndServe(":3000", http.FileServer(http.Dir("")))
		wg.Done()
	}()
	exec.Command("open", "http://localhost:3000/viewer.html").Run()
	wg.Wait()
}
