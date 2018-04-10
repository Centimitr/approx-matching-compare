package main

import (
	"time"
	"os/exec"
	"net/http"
	"sync"
)

func main() {
	t := time.Now()
	am := NewApproxMatchRunner()
	am.Load("task.json")
	am.Stat()
	am.Run(&DirectMatch{}, LIMIT_1)
	am.Run(&NeighbourhoodSearch{K: 1}, LIMIT_1)
	am.Run(&NeighbourhoodSearch{K: 2}, LIMIT_2)
	//am.Run(&NGramDistance{N: 2}, LIMIT_1)
	//am.Run(&NGramDistance{N: 3}, LIMIT_1)
	//am.Run(&NGramDistance{N: 4}, LIMIT_1)
	//am.Run(&GlobalEditDistance{}, LIMIT_1_2)
	am.Save("result.json")
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
