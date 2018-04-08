package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	//am := NewApproxMatchRunner()
	//am.Load("task.json")
	//am.Stat()
	since := time.Since(t)
	since = time.Hour

	fmt.Println("TIME: " + since.String())
}
