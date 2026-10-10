package march

import (
	"time"

	"github.com/DiduLin/slg_ddd_design/types"
)

type MarchTaskMove struct {
	MarchTaskBase
	dest *types.Point
}

func NewMarchTaskMove(march *March, speed time.Duration, dest *types.Point) *MarchTaskMove {
	task := &MarchTaskMove{
		MarchTaskBase: MarchTaskBase{
			march:    march,
			duration: speed,
		},
		dest: dest,
	}

	task.handler = task
	return task
}

func (p *MarchTaskMove) OnStart(now time.Time) {
}

func (p *MarchTaskMove) OnResume(now time.Time) {
}

func (p *MarchTaskMove) OnPause(now time.Time) {
}

func (p *MarchTaskMove) OnComplete(now time.Time) {
	p.march.MoveTo(p.dest.X, p.dest.Y)
}
