package march

import (
	"time"

	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
	"github.com/DiduLin/slg_ddd_design/types"
)

func NewRetreatTask(march *March) *CompositeTask {
	context := &MarchTaskContext{
		MoveContext: MoveContext{
			Command:        enum.MarchCommand_Retreat,
			Speed:          time.Second * 3,
			Start:          march.currPoint,
			End:            &march.homePoint,
			TargetObjectID: march.homeObjectID,
			Path:           nil,
			State:          enum.MarchState_Rout,
		},
	}

	march.taskContext = context

	return NewCompositeTask(
		NewMarchTaskPathfind(march, context),
		NewMarchTaskMove(march, context),
		NewMarchTaskStay(march),
	)
}

func NewMarchStayTask(march *March, speed time.Duration, point *types.Point, path []*types.Point) *CompositeTask {
	context := &MarchTaskContext{
		MoveContext: MoveContext{
			Command: enum.MarchCommand_Stay,
			Speed:   speed,
			Start:   march.currPoint,
			End:     point,
			Path:    path,
			State:   enum.MarchState_Moving,
		},
	}

	march.taskContext = context

	return NewCompositeTask(
		NewMarchTaskMove(march, context),
		NewMarchTaskStay(march),
	)
}
