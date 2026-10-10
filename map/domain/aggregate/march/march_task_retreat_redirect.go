package march

import (
	"time"

	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
	"github.com/DiduLin/slg_ddd_design/types"
)

type MarchTaskRetreatRedirect struct {
	marchTaskBase
	point          types.Point
	targetObjectID int64
}

func NewMarchTaskRetreatRedirect(march *March, point types.Point, targetObjectID int64, completeCallback func(time.Time)) *MarchTaskRetreatRedirect {
	task := &MarchTaskRetreatRedirect{
		marchTaskBase: marchTaskBase{
			march:            march,
			completeCallback: completeCallback,
		},
		point:          point,
		targetObjectID: targetObjectID,
	}
	task.handler = task
	return task
}

func (p *MarchTaskRetreatRedirect) OnStart(now time.Time) {
	p.march.MoveTo(p.point.X, p.point.Y)
	p.march.changeState(enum.MarchState_Idle)
}

func (p *MarchTaskRetreatRedirect) OnComplete(now time.Time) {
}

func (p *MarchTaskRetreatRedirect) OnPause(now time.Time) {
}

func (p *MarchTaskRetreatRedirect) OnResume(now time.Time) {
}
