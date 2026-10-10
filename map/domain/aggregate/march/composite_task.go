package march

import (
	"time"
)

type CompositeTask struct {
	tasks []MarchTask
}

func NewCompositeTask(tasks ...MarchTask) *CompositeTask {
	return &CompositeTask{
		tasks: tasks,
	}
}

func (p *CompositeTask) Start(now time.Time) {
	if len(p.tasks) > 0 {
		p.startTask0(now)
	}
}

func (p *CompositeTask) IsStarted() bool {
	if len(p.tasks) == 0 {
		return true
	}

	return p.tasks[0].IsStarted()
}

func (p *CompositeTask) Deadline() time.Time {
	if len(p.tasks) == 0 {
		return time.Now()
	}

	return p.tasks[0].Deadline()
}

func (p *CompositeTask) Pause(now time.Time) {
	if len(p.tasks) == 0 {
		return
	}

	p.tasks[0].Pause(now)
}

func (p *CompositeTask) IsPaused() bool {
	if len(p.tasks) == 0 {
		return false
	}

	return p.tasks[0].IsPaused()
}

func (p *CompositeTask) Resume(now time.Time) {
	if !p.IsPaused() {
		return
	}

	p.tasks[0].Resume(now)
}

func (p *CompositeTask) Complete(now time.Time) {
	if len(p.tasks) == 0 {
		return
	}

	p.tasks[0].Complete(now)
	if !p.tasks[0].IsComplete() {
		return
	}

	p.tasks = p.tasks[1:]
	if len(p.tasks) > 0 {
		p.startTask0(now)
	}
}

func (p *CompositeTask) IsComplete() bool {
	if len(p.tasks) == 0 {
		return true
	}

	return p.tasks[0].IsComplete()
}

func (p *CompositeTask) Update(now time.Time) {
	if p.IsComplete() {
		return
	}

	if now.Equal(p.Deadline()) || now.After(p.Deadline()) {
		p.Complete(now)
	}
}

func (p *CompositeTask) startTask0(now time.Time) {
	task0 := p.tasks[0]
	task0.Start(now)
}
