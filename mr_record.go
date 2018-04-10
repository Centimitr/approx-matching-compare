package main

type ApproxMatchRecord struct {
	Method     string
	Name       string
	StartTime  int
	Candidates [][]string
	Times      []int

	// stat
	Hits      []bool  `json:",omitempty"`
	Precision float64 `json:",omitempty"`
	Recall    float64 `json:",omitempty"`
	FMeasure  float64 `json:",omitempty"`

	TotalTime int     `json:",omitempty"`
	AvgTime   float64 `json:",omitempty"`
	MinTime   int     `json:",omitempty"`
	MaxTime   int     `json:",omitempty"`
	TimeDrop  int     `json:",omitempty"`

	TimeCmp     float64 `json:",omitempty"`
	MinTimeCmp  float64 `json:",omitempty"`
	MaxTimeCmp  float64 `json:",omitempty"`
	TimeDropCmp float64 `json:",omitempty"`

	Lock bool `json:",omitempty"`
}

func numbers(ns []int) (sum int, avg float64, min, max int) {
	for i, v := range ns {
		sum += v
		switch {
		case i == 0:
			min = v
			max = v
		case v < min:
			min = v
		case v > max:
			max = v
		}
	}
	avg = float64(sum) / float64(len(ns))
	return
}

func (r *ApproxMatchRecord) Stat(corrects []string) {
	if r.Lock {
		return
	}
	r.TotalTime, r.AvgTime, r.MinTime, r.MaxTime = numbers(r.Times)
}

func (r *ApproxMatchRecord) Cmp(baseRecord ApproxMatchRecord) {
	r.TimeCmp = float64(r.AvgTime) / float64(baseRecord.AvgTime)
	r.MinTimeCmp = float64(r.MinTime) / float64(baseRecord.MinTime)
	r.MaxTimeCmp = float64(r.MaxTime) / float64(baseRecord.MaxTime)
	r.TimeDropCmp = float64(r.TimeDrop) / float64(baseRecord.TimeDrop)
}
