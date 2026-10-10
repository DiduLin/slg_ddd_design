package march

import (
	"time"

	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
)

type MarchTaskStay struct {
	marchTaskBase
}

func NewMarchTaskStay(march *March) *MarchTaskStay {
	task := &MarchTaskStay{
		marchTaskBase: marchTaskBase{
			march: march,
		},
	}

	task.handler = task
	return task
}

func (p *MarchTaskStay) OnStart(now time.Time) {
}

func (p *MarchTaskStay) OnComplete(now time.Time) {
	p.march.changeState(enum.MarchState_Stay)
}

func (p *MarchTaskStay) OnPause(now time.Time) {
}

func (p *MarchTaskStay) OnResume(now time.Time) {
	p.march.changeState(enum.MarchState_Stay)
}
