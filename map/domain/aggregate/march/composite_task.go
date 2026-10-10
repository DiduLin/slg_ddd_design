package march

import (
	"time"
)

type CompositeTask struct {
	tasks   []MarchTask
	handler MarchTaskHandler
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

	if p.handler != nil {
		p.handler.OnStart(now)
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

	if p.handler != nil {
		p.handler.OnPause(now)
	}
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

	if p.handler != nil {
		p.handler.OnResume(now)
	}
}

func (p *CompositeTask) Complete(now time.Time) {
	if p.handler != nil {
		p.handler.OnComplete(now)
	}
}

func (p *CompositeTask) IsComplete() bool {
	return len(p.tasks) == 0
}

func (p *CompositeTask) Update(now time.Time) {
	if p.IsComplete() {
		return
	}

	task0 := p.tasks[0]
	task0.Update(now)

	if task0.IsComplete() {
		p.tasks = p.tasks[1:]
		if p.IsComplete() {
			p.Complete(now)
		} else {
			p.startTask0(now)
		}
	}

}

func (p *CompositeTask) startTask0(now time.Time) {
	task0 := p.tasks[0]
	task0.Start(now)
}
