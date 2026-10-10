package march

import (
	"fmt"
	"time"
)

type MarchTaskBattleWait struct {
	MarchTaskBase
}

func NewMarchTaskBattleWait(march *March) *MarchTaskBattleWait {
	task := &MarchTaskBattleWait{
		MarchTaskBase: MarchTaskBase{
			march:    march,
			duration: time.Hour,
		},
	}

	task.handler = task
	return task
}

func (p *MarchTaskBattleWait) OnStart(now time.Time) {
}

func (p *MarchTaskBattleWait) OnComplete(now time.Time) {
	fmt.Printf("%s: Battle complete: %d\n", p.march.clock, p.march.objectID)
}

func (p *MarchTaskBattleWait) OnPause(now time.Time) {
}

func (p *MarchTaskBattleWait) OnResume(now time.Time) {
}
