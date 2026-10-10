package march

import (
	"fmt"
	"slices"
	"time"

	"github.com/DiduLin/slg_ddd_design/map/acl"
	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
	"github.com/DiduLin/slg_ddd_design/map/domain/event"
	"github.com/DiduLin/slg_ddd_design/types"
)

type March struct {
	objectID     int64
	currPoint    *types.Point
	tasks        []MarchTask
	state        enum.MarchState
	eventbus     acl.Eventbus
	clock        acl.Clock
	homePoint    types.Point
	homeObjectID int64

	pathfindGateway acl.PathFindGateway
	taskContext     *MarchTaskContext
}

func NewMarch(objectID int64, eventbus acl.Eventbus, clock acl.Clock, pathfindGateway acl.PathFindGateway) *March {
	return &March{
		objectID:        objectID,
		currPoint:       types.NewPoint(0, 0),
		state:           enum.MarchState_Idle,
		eventbus:        eventbus,
		clock:           clock,
		pathfindGateway: pathfindGateway,
	}
}

func (p *March) Update(now time.Time) {
	for len(p.tasks) > 0 {
		task := p.tasks[len(p.tasks)-1]

		if task.IsComplete() {
			p.tasks = p.tasks[:len(p.tasks)-1]
			continue
		}

		if task.IsPaused() {
			task.Resume(now)
		} else if !task.IsStarted() {
			task.Start(now)
		}

		task.Update(now)

		if !task.IsComplete() {
			break
		}
	}
}

func (p *March) StartMarchStay(speed time.Duration, point *types.Point, path []*types.Point) {
	p.tasks = nil
	p.currPoint = path[0]
	p.addTask(NewMarchStayTask(p, speed, point, path))
}

func (p *March) TrapedInBattle(now time.Time) {
	p.pauseTask(now)
	p.changeState(enum.MarchState_BattleWait)
	p.addTask(NewMarchTaskBattleWait(p))
}

func (p *March) BattleEnd(now time.Time) {
	if p.state != enum.MarchState_BattleWait && p.state != enum.MarchState_Battle {
		return
	}

	if len(p.tasks) == 0 {
		return
	}

	p.tasks[len(p.tasks)-1].Complete(now)
}

func (p *March) BattleFail(now time.Time) {
	p.tasks = nil
	p.addTask(NewRetreatTask(p))
}

func (p *March) changeState(state enum.MarchState) {
	p.state = state
	fmt.Printf("%s: March %d change state: %s\n", p.clock, p.objectID, p.state)
}

// addTask 添加任务
// 任务按添加顺序执行
func (p *March) addTask(tasks ...MarchTask) {
	for _, task := range slices.Backward(tasks) {
		p.tasks = append(p.tasks, task)
	}
}

func (p *March) pauseTask(now time.Time) {
	if len(p.tasks) > 0 {
		p.tasks[len(p.tasks)-1].Pause(now)
	}
}

func (p *March) moveNext() {
	ctx := p.taskContext
	if ctx.PathIndex < len(ctx.Path)-1 {
		ctx.PathIndex++
		point := ctx.Path[ctx.PathIndex]
		p.MoveTo(point.X, point.Y)
	}
}

func (p *March) requestPathFind() error {
	ctx := p.taskContext
	return p.pathfindGateway.FindPath(p.objectID, ctx.Command, ctx.TargetObjectID, ctx.Start, ctx.End)
}

func (p *March) pathfindFail() {
	p.tasks = nil

	// 回城寻路失败，直接返回出生点
	if p.taskContext.Command == enum.MarchCommand_Retreat {
		p.changeState(p.taskContext.State)
		p.addTask(NewMarchTaskRetreatRedirect(p, *p.taskContext.End, p.taskContext.TargetObjectID))
	}
}

func (p *March) RespondPathfind(path []*types.Point) {
	p.taskContext.Path = path
	if len(p.tasks) > 0 {
		task := p.tasks[len(p.tasks)-1]
		if pathfindTask, ok := task.(*MarchTaskPathFind); ok {
			pathfindTask.deadline = p.clock.Now()
		}
	}
}

func (p *March) MoveTo(x, y int32) {
	from := p.currPoint
	p.currPoint = types.NewPoint(x, y)
	fmt.Printf("%s: March %d move from %s to %s\n", p.clock, p.objectID, from, p.currPoint)

	p.eventbus.Publish(&event.MarchMoved{
		ObjectID: p.objectID,
		From:     types.MakePoint(from.X, from.Y),
		To:       types.MakePoint(p.currPoint.X, p.currPoint.Y),
	})
}
