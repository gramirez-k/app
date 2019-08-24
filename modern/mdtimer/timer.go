package mdtimer

import (
	"app/fw"
	"time"
)

type Timer struct{}

func (t Timer) Now() time.Time {
	return time.Now()
}

func NewTimer() fw.Timer {
	return Timer{}
}