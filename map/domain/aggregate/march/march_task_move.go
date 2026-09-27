package march

import (
	"fmt"
	"time"
)

type MarchTaskMove struct {
	marchTaskBase
	context *MarchMoveContext
}

func NewMarchTaskMove(march *March, context *MarchMoveContext, completeCallback func(time.Time)) *MarchTaskMove {
	task := &MarchTaskMove{
		marchTaskBase: marchTaskBase{
			march:            march,
			completeCallback: completeCallback,
			duration:         context.Speed,
		},
		context: context,
	}

	task.handler = task
	return task
}

func (p *MarchTaskMove) OnStart(now time.Time) {
	p.march.changeState(p.context.State)
}

func (p *MarchTaskMove) OnResume(now time.Time) {
	p.march.changeState(p.context.State)
}

func (p *MarchTaskMove) OnPause(now time.Time) {
}

func (p *MarchTaskMove) OnComplete(now time.Time) {
	if p.context.End.X != p.march.currPoint.X || p.context.End.Y != p.march.currPoint.Y {
		p.deadline = now.Add(p.duration)
		fmt.Printf("deadline: %s\n", p.deadline.Format(time.DateTime))

		p.isComplete = false
	}
}

func (p *MarchTaskMove) OnFail(now time.Time) {
}
