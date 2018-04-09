package main

import (
	"time"
)

func main() {
	t := time.Now()
	am := NewApproxMatchRunner()
	am.Load("task.json")
	am.Stat()
	am.Run(DirectMatch{})
	//am.Run(NeighbourhoodSearch{K: 1})
	am.Save("result.json")
	since := time.Since(t)
	print("TIME: " + since.String())
}
