package march

import (
	"fmt"
	"time"
)

type MarchTaskPathFind struct {
	MarchTaskBase
	pathfindContext *MarchTaskContext
}

func NewMarchTaskPathfind(march *March, moveContext *MarchTaskContext) *MarchTaskPathFind {
	task := &MarchTaskPathFind{
		MarchTaskBase: MarchTaskBase{
			march:    march,
			duration: time.Second * 3,
		},
		pathfindContext: moveContext,
	}
	task.handler = task
	return task
}

func (p *MarchTaskPathFind) OnStart(now time.Time) {
	if err := p.march.requestPathFind(); err != nil {
		fmt.Printf("March %d find path fail: %v\v", p.march.objectID, err)
		p.march.pathfindFail()
	}
}

func (p *MarchTaskPathFind) OnComplete(now time.Time) {
	// 路径为空，表示超时失败
	if len(p.pathfindContext.Path) == 0 {
		p.march.pathfindFail()
		return
	}
}

func (p *MarchTaskPathFind) OnPause(now time.Time) {
}

func (p *MarchTaskPathFind) OnResume(now time.Time) {
}
