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
		Fail(now time.Time)
	}

	MarchTaskHandler interface {
		OnStart(now time.Time)
		OnComplete(now time.Time)
		OnPause(now time.Time)
		OnResume(now time.Time)
		OnFail(now time.Time)
	}

	marchTaskBase struct {
		started          bool
		duration         time.Duration
		deadline         time.Time
		paused           bool
		pauseTime        time.Time
		march            *March
		completeCallback func(time.Time)
		isComplete       bool
		handler          MarchTaskHandler
	}
)

func (p *marchTaskBase) Start(now time.Time) {
	p.started = true
	p.deadline = now.Add(p.duration)

	fmt.Printf("deadline: %s\n", p.deadline.Format(time.DateTime))

	if p.handler != nil {
		p.handler.OnStart(now)
	}
}

func (p *marchTaskBase) IsStarted() bool {
	return p.started
}

func (p *marchTaskBase) Deadline() time.Time {
	return p.deadline
}

func (p *marchTaskBase) Pause(now time.Time) {
	p.paused = true
	p.pauseTime = now

	if p.handler != nil {
		p.handler.OnPause(now)
	}
}

func (p *marchTaskBase) IsPaused() bool {
	return p.paused
}

func (p *marchTaskBase) Resume(now time.Time) {
	if !p.IsPaused() {
		return
	}

	p.paused = false
	p.deadline = p.deadline.Add(now.Sub(p.pauseTime))

	if p.handler != nil {
		p.handler.OnResume(now)
	}
}

func (p *marchTaskBase) Complete(now time.Time) {
	if p.completeCallback != nil {
		p.completeCallback(now)
	}

	p.isComplete = true

	if p.handler != nil {
		p.handler.OnComplete(now)
	}
}

func (p *marchTaskBase) IsComplete() bool {
	return p.isComplete
}

func (p *marchTaskBase) Update(now time.Time) {
	if p.IsComplete() {
		return
	}

	if now.Equal(p.deadline) || now.After(p.deadline) {
		p.Complete(now)
	}
}

func (p *marchTaskBase) Fail(now time.Time) {
	if p.isComplete {
		return
	}

	if p.handler != nil {
		p.handler.OnFail(now)
	}
}
