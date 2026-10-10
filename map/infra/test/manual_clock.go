package test

import "time"

type ManualClock struct {
	t time.Time
}

func NewManualClock(t time.Time) *ManualClock {
	return &ManualClock{t: t}
}

func (p *ManualClock) Now() time.Time {
	return p.t
}

func (p *ManualClock) After(d time.Duration) {
	p.t = p.t.Add(d)
}

func (p *ManualClock) String() string {
	return p.t.Format("2006-01-02 15:04:05")
}
