package main

// util
func minDistance(list []string, handle func(t string) int) string {
	minD := 0
	minI := 0
	for i, t := range list {
		d := handle(t)
		if i == 0 {
			minD = d
		} else if d < minD {
			minD = d
			minI = i
		}
	}
	return list[minI]
}

