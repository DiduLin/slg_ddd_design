package march

import (
	"time"

	"github.com/DiduLin/slg_ddd_design/map/domain/enum"
	"github.com/DiduLin/slg_ddd_design/types"
)

type (
	MarchMoveContext struct {
		Command        enum.MarchCommand // 造成行军移动的指令
		Start          *types.Point      // 移动的起点
		End            *types.Point      // 移动的终点
		TargetObjectID int64             // 移动的目标对象ID
		Path           []*types.Point    // 移动的路径：可以由外部传入（主动发起的移动），或者由寻路算法生成（被动触发的移动）
		State          enum.MarchState   // 移动时的状态
		Speed          time.Duration     // 移动时的速度（走一格所需要的时间）
		PathIndex      int               // 当前路径索引
	}
)
