package main

import (
	"fmt"
	"time"
)

type Time = time.Time
type Duration = time.Duration

type Chronometer struct {
	start  Time
	finish Time
}

func (chronometer *Chronometer) Start() {
	chronometer.start = time.Now()
}

func (chronometer *Chronometer) GetEllapsedTime() string {
	ellapsedTime := time.Since(chronometer.start)

	total := int(ellapsedTime.Seconds())
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total) % 60

	chronoString := fmt.Sprintf("%02d:%02d:%02d\n", hours, minutes, seconds)
	return chronoString
}

func (chronometer *Chronometer) Stop() {
	chronometer.finish = time.Now()
}
