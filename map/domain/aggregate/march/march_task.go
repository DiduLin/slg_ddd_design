package march

import (
	"fmt"
	"time"
)

type (
	MarchTask interface {
		Start(now time.Time)
		IsStarted() bool
		Deadline() time.Time
		Pause(time.Time)
		IsPaused() bool
		Resume(time.Time)
		Complete(time.Time)
		IsComplete() bool
		Update(now time.Time)
	}

	MarchTaskHandler interface {
		OnStart(now time.Time)
		OnComplete(now time.Time)
		OnPause(now time.Time)
		OnResume(now time.Time)
	}

	MarchTaskBase struct {
		started    bool
		duration   time.Duration
		deadline   time.Time
		paused     bool
		pauseTime  time.Time
		march      *March
		isComplete bool
		handler    MarchTaskHandler
	}
)

func (p *MarchTaskBase) Start(now time.Time) {
	p.started = true
	p.deadline = now.Add(p.duration)

	fmt.Printf("deadline: %s\n", p.deadline.Format(time.DateTime))

	if p.handler != nil {
		p.handler.OnStart(now)
	}
}

func (p *MarchTaskBase) IsStarted() bool {
	return p.started
}

func (p *MarchTaskBase) Deadline() time.Time {
	return p.deadline
}

func (p *MarchTaskBase) Pause(now time.Time) {
	p.paused = true
	p.pauseTime = now

	if p.handler != nil {
		p.handler.OnPause(now)
	}
}

func (p *MarchTaskBase) IsPaused() bool {
	return p.paused
}

func (p *MarchTaskBase) Resume(now time.Time) {
	if !p.IsPaused() {
		return
	}

	p.paused = false
	p.deadline = p.deadline.Add(now.Sub(p.pauseTime))

	if p.handler != nil {
		p.handler.OnResume(now)
	}
}

func (p *MarchTaskBase) Complete(now time.Time) {
	p.isComplete = true

	if p.handler != nil {
		p.handler.OnComplete(now)
	}
}

func (p *MarchTaskBase) IsComplete() bool {
	return p.isComplete
}

func (p *MarchTaskBase) Update(now time.Time) {
	if p.IsComplete() {
		return
	}

	if now.Equal(p.deadline) || now.After(p.deadline) {
		p.Complete(now)
	}
}
