package march

import "time"

type MarchTaskMoveChain struct {
	*CompositeTask
	march   *March
	context *MarchTaskContext
}

func NewMarchTaskMoveChain(march *March, context *MarchTaskContext) *MarchTaskMoveChain {
	var tasks []MarchTask
	for i := 1; i < len(context.Path); i++ {
		tasks = append(tasks, NewMarchTaskMove(march, context.Speed, context.Path[i]))
	}

	task := &MarchTaskMoveChain{
		CompositeTask: NewCompositeTask(tasks...),
		march:         march,
		context:       context,
	}

	task.handler = task
	return task
}

func (p *MarchTaskMoveChain) OnStart(now time.Time) {
	p.march.changeState(p.context.State)
}

func (p *MarchTaskMoveChain) OnComplete(now time.Time) {
}

func (p *MarchTaskMoveChain) OnPause(now time.Time) {
}

func (p *MarchTaskMoveChain) OnResume(now time.Time) {
	p.march.changeState(p.context.State)
}
