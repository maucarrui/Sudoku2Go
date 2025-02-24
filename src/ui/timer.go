package ui

import (
	"math"
	"strconv"
	"time"
)

type Time = time.Time

type Timer struct {
	initialTime Time
}

func NewTimer() *Timer {
	timer := &Timer{
		initialTime: time.Now(),
	}

	return timer
}

func (timer *Timer) ToString() string {
	initialTime := timer.initialTime
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(initialTime)

	elapsedSeconds := elapsedTime.Seconds()

	seconds := int(math.Mod(elapsedSeconds, 60))
	minutes := int(math.Mod(elapsedSeconds/60, 60))
	hours := int((elapsedSeconds / 60) / 60)

	ss := ""
	mm := ""
	hh := ""

	if seconds < 10 {
		ss += "0"
	}
	ss += strconv.Itoa(seconds)

	if minutes < 10 {
		mm += "0"
	}
	mm += strconv.Itoa(minutes)

	if hours < 10 {
		hh += "0"
	}
	hh += strconv.Itoa(hours)

	return hh + ":" + mm + ":" + ss
}
